package algorithm

import (
	"sort"
	"strconv"
)

// ThreeSum returns all the "unique" triplets that have their sums to be 0.
// LeetCode #15
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	size := len(nums)
	if size < 3 {
		return res
	}

	hasPositive := false
	countOfNegatives := 0
	countOfZero := 0
	for _, v := range nums {
		if v > 0 {
			hasPositive = true
		} else if v < 0 {
			countOfNegatives++
		} else if countOfZero < 3 {
			countOfZero++
		}
	}
	if countOfZero >= 3 {
		res = append(res, []int{0, 0, 0})
	}
	if countOfNegatives == 0 || !hasPositive {
		return res
	}

	sort.Ints(nums)

	left := 0
	right := size - 1
	prevPos := -1
	prevNeg := 1
	moveState := 0
	for left < countOfNegatives && right >= countOfNegatives {
		negative := nums[left]
		positive := nums[right]
		if moveState == 1 && negative == prevNeg {
			left++
			continue
		} else if moveState == 2 && positive == prevPos {
			right--
			continue
		}

		if positive+negative > 0 {
			// it's a neg-neg-pos trio
			idxNegLeft := left
			idxNegRight := countOfNegatives - 1
			prevNegLeft := 1
			prevNegRight := 1
			lastRoundSuccessful := false
			for idxNegLeft < idxNegRight {
				valNegLeft := nums[idxNegLeft]
				valNegRight := nums[idxNegRight]
				if lastRoundSuccessful && valNegLeft == prevNegLeft && valNegRight == prevNegRight {
					idxNegLeft++
					idxNegRight--
					continue
				}
				lastRoundSuccessful = false
				sum := valNegLeft + valNegRight + positive
				if sum > 0 {
					idxNegRight--
				} else if sum < 0 {
					idxNegLeft++
				} else {
					lastRoundSuccessful = true
					res = append(res, []int{valNegLeft, valNegRight, positive})
					idxNegLeft++
					idxNegRight--
				}
				prevNegLeft = valNegLeft
				prevNegRight = valNegRight
			}
			prevPos = positive
			moveState = 2
			right--
		} else {
			// it's a neg-zero-pos or neg-pos-pos trio
			idxPosLeft := countOfNegatives
			idxPosRight := right
			prevPosLeft := -1
			prevPosRight := -1
			lastRoundSuccessful := false
			for idxPosLeft < idxPosRight {
				valPosLeft := nums[idxPosLeft]
				valPosRight := nums[idxPosRight]
				if lastRoundSuccessful && prevPosLeft == valPosLeft && prevPosRight == valPosRight {
					idxPosLeft++
					idxPosRight--
					continue
				}
				lastRoundSuccessful = false
				sum := valPosLeft + valPosRight + negative
				if sum > 0 {
					idxPosRight--
				} else if sum < 0 {
					idxPosLeft++
				} else {
					lastRoundSuccessful = true
					res = append(res, []int{valPosLeft, valPosRight, negative})
					idxPosLeft++
					idxPosRight--
				}
				prevPosLeft = valPosLeft
				prevPosRight = valPosRight
			}
			prevNeg = negative
			moveState = 1
			left++
		}
	}

	return res
}

func threeSumHashVersion(nums []int) [][]int {
	res := [][]int{}

	if len(nums) < 3 {
		return res
	}

	hasPositive := false
	hasNegative := false
	countOfZero := 0

	for _, v := range nums {
		if v > 0 {
			hasPositive = true
		} else if v < 0 {
			hasNegative = true
		} else if countOfZero < 3 {
			countOfZero++
		}
	}

	if !hasNegative || !hasPositive {
		if countOfZero >= 3 {
			res = append(res, []int{0, 0, 0})
		}
		return res
	}

	resMap := map[string]bool{}
	m := map[int][]int{}

	for i, v := range nums {
		if cur, exists := m[v]; exists {
			updated := append(cur, i)
			m[v] = updated
		} else {
			cur = []int{i}
			m[v] = cur
		}
	}

	for idx1, v1 := range nums {
		for j, v2 := range nums[idx1+1:] {
			idx2 := idx1 + 1 + j
			target := -1 * (v1 + v2)
			if idxList, exists := m[target]; exists {
				for _, idx3 := range idxList {
					if idx1 != idx3 && idx2 != idx3 {
						a, b, c := ascendThreeIntegers(v1, v2, target)
						resKey := strconv.Itoa(a) + "," + strconv.Itoa(b) + "," + strconv.Itoa(c)
						if !resMap[resKey] {
							resMap[resKey] = true
							res = append(res, []int{v1, v2, target})
						}
					}
				}
			}
		}
	}

	return res
}

func ascendThreeIntegers(a, b, c int) (int, int, int) {
	if a <= b {
		if a <= c {
			if b <= c {
				return a, b, c
			} else {
				return a, c, b
			}
		} else {
			return c, a, b
		}
	} else {
		if a <= c {
			return b, a, c
		} else {
			if b <= c {
				return b, c, a
			} else {
				return c, b, a
			}
		}
	}
}
