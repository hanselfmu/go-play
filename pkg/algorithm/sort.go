package algorithm

import "fmt"

// QuickSortBasic implements the quicksort algorithm
// worst case theta(n^2), average case expected O(nlogn)
// This implementation is in-place
func QuickSortBasic(A []int) []int {
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
	result = append(result, QuickSortBasic(left)...)
	result = append(result, pivotVal)
	result = append(result, QuickSortBasic(right)...)
	return result
}

// HeapSortBasic implements a simple heapsort with an array as heap
// This implementation is in-place.
func HeapSortBasic(A []int) []int {
	size := len(A)
	if size <= 1 {
		return A
	}
	// Min Heap is a full tree, so we don't need a binary linked list to represent;
	// a dynamic array (like a slice in Go) will suffice.

	// build up a min heap, O(nlogn)
	for idx, val := range A[1:] {
		curIdx := idx + 1
		parentIdx := (curIdx - 1) / 2
		A[curIdx] = val
		curVal := val
		parentVal := A[parentIdx]
		fmt.Printf("cur %d; parent %d\n", curIdx, parentIdx)
		fmt.Printf("now: %v\n", A)
		for curVal < parentVal {
			A[curIdx] = parentVal
			A[parentIdx] = curVal
			curIdx = parentIdx
			parentIdx = (curIdx - 1) / 2
			curVal = A[curIdx]
			parentVal = A[parentIdx]
		}
	}

	// pending siftDowns

	return A
}
