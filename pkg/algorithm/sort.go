package algorithm

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
// This implementation also DELIBERATELY refrains from using subroutines. I know it saves a lot of code and makes the structure cleaner, but I just like copy-pasting more right now.
func HeapSortBasic(A []int) []int {
	size := len(A)
	if size <= 1 {
		return A
	}
	// Min Heap is a complete tree, so we don't need a binary linked list to represent;
	// a dynamic array (like a slice in Go) will suffice.

	// build up a min heap, O(nlogn)
	for idx, val := range A[1:] {
		curIdx := idx + 1
		parentIdx := (curIdx - 1) / 2
		A[curIdx] = val
		curVal := val
		parentVal := A[parentIdx]
		for curVal < parentVal {
			A[curIdx] = parentVal
			A[parentIdx] = curVal
			curIdx = parentIdx
			parentIdx = (curIdx - 1) / 2
			curVal = A[curIdx]
			parentVal = A[parentIdx]
		}
	}

	// doing siftDowns from the bottom of the tree by:
	// 1. moving the topmost element (currently smallest) to the end of the current sorted list;
	// 2. taking the last element in the array and put it at the now-empty root;
	// 3. move this newly placed element down the tree until the minHeap property is restored
	sortedStart := size - 1
	for sortedStart >= 0 {
		curMin := A[0]
		A[0] = A[sortedStart]
		A = append(append(A[0:sortedStart], A[sortedStart+1:]...), curMin)

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
				if curVal > leftVal && curVal > rightVal {
					if leftVal > rightVal {
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
				} else if curVal > leftVal {
					temp := A[curIdx]
					A[curIdx] = A[leftIdx]
					A[leftIdx] = temp
					curIdx = leftIdx
				} else if curVal > rightVal {
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
				if curVal > leftVal {
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
				if curVal > rightVal {
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
