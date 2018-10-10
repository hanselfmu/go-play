package algorithm

// QuickSort implements the quicksort algorithm
// worst case theta(n^2), average case expected O(nlogn)
func QuickSort(A []int) []int {
	size := len(A)
	if size <= 1 {
		return A
	}
	pivotIdx := size / 2
	pivotVal := A[pivotIdx]
	var left, right, result []int
	for idx, v := range A {
		if pivotIdx != idx {
			if v <= pivotVal {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
	}
	result = append(result, QuickSort(left)...)
	result = append(result, pivotVal)
	result = append(result, QuickSort(right)...)
	return result
}
