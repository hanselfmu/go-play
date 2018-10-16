package algorithm

import "fmt"

// PainterPartition solves the painter's partition problem at
// https://www.geeksforgeeks.org/painters-partition-problem-set-2/
func PainterPartition(k int, seq []int) []int {
	seqLen := len(seq)
	if k == 0 || k == 1 {
		return []int{seqLen}
	}

	partitions := make([]partition, k)

	if k >= seqLen {
		// O(N)
		res := make([]int, seqLen)
		for i := range res {
			res[i] = 1
		}
		return res
	}

	// initialize
	// O(K)
	for i := 0; i < k-1; i++ {
		initialPartition := seq[i]
		partitions[i] = partition{[]int{initialPartition}, initialPartition}
	}
	// O(N)
	remainingSum := calculateSum(seq[k-1:])
	partitions[k-1] = partition{seq[k-1:], remainingSum}

	isStable := false
	// loop until stable condition
	// O(N)
	for !isStable {
		updated := false

		// O(K)
		for i := k - 1; i > 0; i-- {
			curPart := partitions[i]
			prevPart := partitions[i-1]
			curPartLen := len(curPart.values)
			if curPartLen > 1 {
				move := curPart.values[0]
				if getMax(curPart.sum-move, prevPart.sum+move) <= getMax(curPart.sum, prevPart.sum) {
					curPart.values = curPart.values[1:]
					curPart.sum = curPart.sum - move
					prevPart.values = append(prevPart.values, move)
					prevPart.sum = prevPart.sum + move
					partitions[i] = curPart
					partitions[i-1] = prevPart
					updated = true
				}
			}
		}
		isStable = !updated
	}

	// return partition counts
	res := make([]int, k)
	for i := range res {
		res[i] = len(partitions[i].values)
	}
	return res
}

// LargestSumOfAverages calculates the largest sum of partitioning a group of integers
// LeetCode #813
func LargestSumOfAverages(A []int, K int) float64 {
	seq := A
	k := K
	seqLen := len(seq)
	if k == 0 || k == 1 {
		return float64(calculateSum(seq)) / float64(seqLen)
	}

	partitions := make([]partition, k)

	if k >= seqLen {
		// O(N)
		return float64(calculateSum(seq))
	}

	// initialize
	// O(K)
	for i := 0; i < k-1; i++ {
		initialPartition := seq[i]
		partitions[i] = partition{[]int{initialPartition}, initialPartition}
	}
	// O(N)
	remainingSum := calculateSum(seq[k-1:])
	partitions[k-1] = partition{seq[k-1:], remainingSum}

	isStable := false
	// loop until stable condition
	// O(N)
	for !isStable {
		updated := false
		for i, p := range partitions {
			fmt.Printf("partition %d: %v\n", i, p.values)
		}
		// O(K)
		for i := k - 1; i > 0; i-- {
			curPart := partitions[i]
			prevPart := partitions[i-1]
			curPartLen := len(curPart.values)
			prevPartLen := len(prevPart.values)
			if curPartLen > 1 {
				move := curPart.values[0]
				curAvgSum := float64(curPart.sum)/float64(curPartLen) + float64(prevPart.sum)/float64(prevPartLen)
				newAvgSum := float64(curPart.sum-move)/float64(curPartLen-1) + float64(prevPart.sum+move)/float64(prevPartLen+1)

				if newAvgSum >= curAvgSum {
					curPart.values = curPart.values[1:]
					curPart.sum = curPart.sum - move
					prevPart.values = append(prevPart.values, move)
					prevPart.sum = prevPart.sum + move
					partitions[i] = curPart
					partitions[i-1] = prevPart
					updated = true
				}
			}
		}
		isStable = !updated
	}

	sum := 0.0
	for _, part := range partitions {
		sum += float64(part.sum) / float64(len(part.values))
	}
	return sum
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateSum(list []int) int {
	res := 0
	for _, v := range list {
		res += v
	}
	return res
}

type partition struct {
	values []int
	sum    int
}
