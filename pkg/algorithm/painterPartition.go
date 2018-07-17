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
		res := make([]int, seqLen)
		for i := range res {
			res[i] = 1
		}
		return res
	}

	// 3, 2, 5, 7, 4, 2, 9, 3, 8
	// 3 | 2 | 5 7 2 9 3 8
	// initialize
	for i := 0; i < k-1; i++ {
		initialPartition := seq[i]
		partitions[i] = partition{[]int{initialPartition}, initialPartition}
	}
	remainingSum := calculateSum(seq[k:])
	// totalSum := calculateSum(seq)
	// globalMin := totalSum / k
	partitions[k-1] = partition{seq[k:], remainingSum}

	isStable := false
	// loop until stable condition
	for !isStable {
		updated := false

		for i := k - 1; i > 0; i-- {
			curPart := partitions[i]
			prevPart := partitions[i-1]
			curPartLen := len(curPart.values)
			if curPartLen > 1 {
				move := curPart.values[curPartLen-1]
				if getMax(curPart.sum-move, prevPart.sum+move) <= getMax(curPart.sum, prevPart.sum) {
					curPart.values = curPart.values[:curPartLen-1]
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
		fmt.Printf("partition %d: %v with sum %d\n", i, partitions[i].values, partitions[i].sum)
		res[i] = len(partitions[i].values)
	}
	return res
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
