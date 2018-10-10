/*
$ go test ./pkg/algorithm/ -v
*/
package algorithm

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	unsorted := []int{10, -5, 6, 2, -7, 20, 3, 4, -1, 0, 0, 3, 12, 35, 9}
	expected := []int{-7, -5, -1, 0, 0, 2, 3, 3, 4, 6, 9, 10, 12, 20, 35}
	result := QuickSort(unsorted)

	for idx, val := range expected {
		if val != result[idx] {
			t.Errorf("Expected %d at idx %d; got %d\n", val, idx, result[idx])
		}
	}
}
