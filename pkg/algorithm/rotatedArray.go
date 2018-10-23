package algorithm

// SearchRotatedSortedArray searches the index of target in a given sorted array; returns -1 if not found. Needs to be in O(logn) time complexity.
// LeetCode #33
func SearchRotatedSortedArray(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	if size == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if size == 2 {
		if nums[0] == target {
			return 0
		}
		if nums[1] == target {
			return 1
		}
		return -1
	}
	left := nums[0]
	midIdx := (size - 1) / 2
	mid := nums[midIdx]
	right := nums[size-1]
	if left < right {
		// does not have pivot point
		if target < left || target > right {
			return -1
		}
		if target < mid {
			return SearchRotatedSortedArray(nums[:midIdx], target)
		}
		rightHalf := SearchRotatedSortedArray(nums[midIdx:], target)
		if rightHalf != -1 {
			return midIdx + rightHalf
		}
		return -1
	}
	// has pivot point
	leftHalf := SearchRotatedSortedArray(nums[:midIdx], target)
	rightHalf := SearchRotatedSortedArray(nums[midIdx:], target)
	if leftHalf != -1 {
		return leftHalf
	}
	if rightHalf != -1 {
		return midIdx + rightHalf
	}
	return -1
}
