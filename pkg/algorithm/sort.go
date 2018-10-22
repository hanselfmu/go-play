package algorithm

import (
	"runtime"
)

// QuickSortBasic implements the quicksort algorithm
// worst case theta(n^2), average case expected O(nlogn)
// This implementation is out-of-place, but an in-place quicksort is available.
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
// This implementation also DELIBERATELY refrains from using subroutines. I know it saves a lot of code and makes the structure cleaner, but I just like copy-pasting more right now.
func HeapSortBasic(A []int) []int {
	size := len(A)
	if size <= 1 {
		return A
	}
	// Max Heap is a complete tree, so we don't need a binary linked list to represent;
	// a dynamic array (like a slice in Go) will suffice.

	// build up a max heap, O(nlogn)
	for idx, val := range A[1:] {
		curIdx := idx + 1
		parentIdx := (curIdx - 1) / 2
		A[curIdx] = val
		curVal := val
		parentVal := A[parentIdx]
		for curVal > parentVal {
			A[curIdx] = parentVal
			A[parentIdx] = curVal
			curIdx = parentIdx
			parentIdx = (curIdx - 1) / 2
			curVal = A[curIdx]
			parentVal = A[parentIdx]
		}
	}

	// doing siftDowns from the bottom of the tree by:
	// 1. moving the topmost element (currently largest) to the end of the current sorted list;
	// 2. taking the last element in the array and put it at the now-empty root;
	// 3. move this newly placed element down the tree until the maxHeap property is restored
	sortedStart := size - 1
	for sortedStart >= 0 {
		curMax := A[0]
		A[0] = A[sortedStart]
		A[sortedStart] = curMax

		// step 3
		curIdx := 0
		leftIdx := curIdx*2 + 1
		rightIdx := curIdx*2 + 2
		isStable := false
		for !isStable {
			if leftIdx < sortedStart && rightIdx < sortedStart {
				curVal := A[curIdx]
				leftVal := A[leftIdx]
				rightVal := A[rightIdx]
				if curVal < leftVal && curVal < rightVal {
					if leftVal < rightVal {
						// switch cur with right
						temp := A[curIdx]
						A[curIdx] = A[rightIdx]
						A[rightIdx] = temp
						curIdx = rightIdx
					} else {
						// switch cur with left
						temp := A[curIdx]
						A[curIdx] = A[leftIdx]
						A[leftIdx] = temp
						curIdx = leftIdx
					}
				} else if curVal < leftVal {
					temp := A[curIdx]
					A[curIdx] = A[leftIdx]
					A[leftIdx] = temp
					curIdx = leftIdx
				} else if curVal < rightVal {
					temp := A[curIdx]
					A[curIdx] = A[rightIdx]
					A[rightIdx] = temp
					curIdx = rightIdx
				} else {
					isStable = true
				}
			} else if leftIdx < sortedStart {
				curVal := A[curIdx]
				leftVal := A[leftIdx]
				if curVal < leftVal {
					temp := A[curIdx]
					A[curIdx] = A[leftIdx]
					A[leftIdx] = temp
					curIdx = leftIdx
				} else {
					isStable = true
				}
			} else if rightIdx < sortedStart {
				curVal := A[curIdx]
				rightVal := A[rightIdx]
				if curVal < rightVal {
					temp := A[curIdx]
					A[curIdx] = A[rightIdx]
					A[rightIdx] = temp
					curIdx = rightIdx
				} else {
					isStable = true
				}
			} else {
				// reached the bottom
				isStable = true
			}

			leftIdx = curIdx*2 + 1
			rightIdx = curIdx*2 + 2
		}

		sortedStart--
	}

	return A
}

// MergeSortBasic sorts an array of ints using mergesort in O(nlogn).
// This implementation is out-of-place. Note that a carefully designed in-place mergesort is available.
func MergeSortBasic(A []int) []int {
	size := len(A)
	if size <= 1 {
		return A
	}
	mid := size / 2
	left := MergeSortBasic(A[:mid])
	right := MergeSortBasic(A[mid:])

	sizeL := len(left)
	sizeR := len(right)
	idxL := 0
	idxR := 0
	result := make([]int, 0, sizeL+sizeR)
	for idxL < sizeL && idxR < sizeR {
		if left[idxL] < right[idxR] {
			result = append(result, left[idxL])
			idxL++
		} else {
			result = append(result, right[idxR])
			idxR++
		}
	}
	if idxL < sizeL {
		result = append(result, left[idxL:]...)
	} else {
		result = append(result, right[idxR:]...)
	}
	return result
}

// MergeSortGoroutine sorts an array of ints using mergesort with the help of Goroutine.
func MergeSortGoroutine(A []int) []int {
	total := len(A)
	numOfCores := runtime.GOMAXPROCS(0)
	goroutineThreshold := total * 2 / numOfCores
	c := make(chan []int)
	go doMergeSort(A, c, goroutineThreshold)
	return <-c
}

func doMergeSort(A []int, c chan []int, goroutineThreshold int) {
	size := len(A)
	if size <= 1 {
		c <- A
		return
	} else if size == 2 {
		if A[0] < A[1] {
			c <- A
			return
		}
		t := A[1]
		A[1] = A[0]
		A[0] = t
		c <- A
		return

	}
	mid := size / 2
	var left, right []int
	if size > goroutineThreshold {
		// https://stackoverflow.com/questions/25860633/order-of-goroutine-unblocking-on-single-channel
		// as to why a new channel is needed: the unblocking order is not guaranteed, and definitely not LIFO right now.
		newC := make(chan []int)
		go doMergeSort(A[:mid], newC, goroutineThreshold)
		go doMergeSort(A[mid:], newC, goroutineThreshold)
		left, right = <-newC, <-newC
	} else {
		left = MergeSortBasic(A[:mid])
		right = MergeSortBasic(A[mid:])
	}

	sizeL := len(left)
	sizeR := len(right)
	idxL := 0
	idxR := 0
	result := make([]int, 0, sizeL+sizeR)
	for idxL < sizeL && idxR < sizeR {
		if left[idxL] < right[idxR] {
			result = append(result, left[idxL])
			idxL++
		} else {
			result = append(result, right[idxR])
			idxR++
		}
	}
	if idxL < sizeL {
		result = append(result, left[idxL:]...)
	} else {
		result = append(result, right[idxR:]...)
	}
	c <- result
}

// MergeSortIterative implements an iterative version of mergesort.
func MergeSortIterative(A []int) []int {
	size := len(A)
	if size <= 1 {
		return A
	}
	mergeSize := 2
	result := make([]int, size)
	for mergeSize < size*2 {
		start := 0
		end := mergeSize
		if end > size {
			end = size
		}
		blockSize := mergeSize / 2
		mid := start + blockSize
		if mid > size {
			mid = size
		}
		isEndReached := false

		for !isEndReached {
			// merge
			idxTarget := start
			idxL := start
			idxR := mid

			for idxL < mid && idxR < end {
				if A[idxL] < A[idxR] {
					result[idxTarget] = A[idxL]
					idxL++
				} else {
					result[idxTarget] = A[idxR]
					idxR++
				}
				idxTarget++
			}
			for idxL < mid {
				result[idxTarget] = A[idxL]
				idxL++
				idxTarget++
			}
			for idxR < end {
				result[idxTarget] = A[idxR]
				idxR++
				idxTarget++
			}

			start = end
			if end >= size {
				isEndReached = true
			} else {
				start = end
				end += mergeSize
				if end > size {
					end = size
				}
				mid = start + blockSize
				if mid > size {
					mid = size
				}
			}
		}
		mergeSize *= 2
		A = result
		result = make([]int, size)
	}
	return A
}
