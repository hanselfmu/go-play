package algorithm

// KthOfSortedArrays returns the kth element of two sorted arrays, in O(log(m+n)) complexity
func KthOfSortedArrays(arr1, arr2 []int, k int) int {
	len1, len2 := len(arr1), len(arr2)

	// edge cases
	if k >= len1+len2 {
		return -1 // error: k should never be larger than len1+len2
	}
	if len1 == 0 {
		return arr2[k]
	}
	if len2 == 0 {
		return arr1[k]
	}
	if k == 0 {
		if arr1[0] < arr2[0] {
			return arr1[0]
		}
		return arr2[0]
	}

	// base case
	if k == 1 {
		if arr1[0] < arr2[0] {
			if len1 > 1 && arr1[1] < arr2[0] {
				return arr1[1]
			}
			return arr2[0]
		}

		if len2 > 1 && arr2[1] < arr1[0] {
			return arr2[1]
		}

		return arr1[0]
	}

	split := k / 2
	if split >= len1 {
		split = len1 - 1
	}
	if split >= len2 {
		split = len2 - 1
	}

	if arr1[split] < arr2[split] {
		if split == 0 {
			split++
		}
		return KthOfSortedArrays(arr1[split:len1], arr2, k-split)
	} else {
		if split == 0 {
			split++
		}
		return KthOfSortedArrays(arr1, arr2[split:len2], k-split)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		medianLeft, medianRight := KthOfSortedArrays(nums1, nums2, totalLen/2-1), KthOfSortedArrays(nums1, nums2, totalLen/2)
		return float64(medianLeft+medianRight) / 2
	}
	return float64(KthOfSortedArrays(nums1, nums2, totalLen/2))
}
