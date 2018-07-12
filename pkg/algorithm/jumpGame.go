package algorithm

// CanJump is from LeetCode #55
// Given an array of non-negative integers,
// you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Determine if you are able to reach the last index.
func CanJump(nums []int) bool {
	target := len(nums) - 1
	if target == -1 {
		return false
	}
	if target == 0 {
		return true
	}

	max := 0
	for idx, cur := range nums {
		curMax := idx + cur
		if curMax > max {
			max = curMax
		}

		if idx >= max {
			return false
		}

		if max >= target {
			return true
		}
	}

	return false
}

// MinJump is from LeetCode #45
// Given an array of non-negative integers,
// you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Calculate the minimum jumps required to reach (or exceed) the last index.
// This solution uses a simple DP
func MinJump(nums []int) int {
	target := len(nums) - 1
	if target <= 0 {
		return 0
	}

	minSteps := make([]int, len(nums))
	max := 0
	for idx, cur := range nums {
		curMax := idx + cur
		if curMax > max {
			max = curMax
		}

		if idx >= max {
			return minSteps[target]
		}

		nextStepCount := minSteps[idx] + 1
		for jumpIdx := idx + 1; jumpIdx <= curMax && jumpIdx <= target; jumpIdx++ {
			curMin := minSteps[jumpIdx]
			if nextStepCount < curMin || curMin == 0 /* initial value */ {
				minSteps[jumpIdx] = nextStepCount
			}
		}

		if curMax >= target && idx < target {
			return nextStepCount
		}
	}

	return minSteps[target]
}

// MinJump2 is an optimized version of MinJump, using Greedy algorithm.
func MinJump2(nums []int) int {
	target := len(nums) - 1
	if target <= 0 {
		return 0
	}

	startIdx := 0
	jumps := 1
	max := nums[0]
	for startIdx < len(nums) {
		if max >= target {
			return jumps
		}
		oldMax := max
		max = 0
		for i := startIdx; i <= oldMax; i++ {
			if nums[i]+i > max {
				max = nums[i] + i
			}
		}
		startIdx = oldMax + 1
		jumps++
	}

	return jumps
}
