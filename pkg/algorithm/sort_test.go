/*
$ go test ./pkg/algorithm/ -v
$ go test ./pkg/algorithm/ -bench=Sort -v
*/
package algorithm

import (
	"math/rand"
	"testing"
)

func TestQuickSortBasic(t *testing.T) {
	unsorted := []int{10, -5, 6, 2, -7, 20, 3, 4, -1, 0, 0, 3, 12, 35, 9}
	expected := []int{-7, -5, -1, 0, 0, 2, 3, 3, 4, 6, 9, 10, 12, 20, 35}
	result := QuickSortBasic(unsorted)

	for idx, val := range expected {
		if val != result[idx] {
			t.Errorf("QuickSortBasic -- Expected %d at idx %d; got %d\n", val, idx, result[idx])
		}
	}
}

func TestHeapSortBasic(t *testing.T) {
	unsorted := []int{10, -5, 6, 2, -7, 20, 3, 4, -1, 0, 0, 3, 12, 35, 9}
	expected := []int{-7, -5, -1, 0, 0, 2, 3, 3, 4, 6, 9, 10, 12, 20, 35}
	result := HeapSortBasic(unsorted)

	for idx, val := range expected {
		if val != result[idx] {
			t.Errorf("HeapSortBasic -- Expected %d at idx %d; got %d\n", val, idx, result[idx])
		}
	}
}

func TestMergeSortBasic(t *testing.T) {
	unsorted := []int{10, -5, 6, 2, -7, 20, 3, 4, -1, 0, 0, 3, 12, 35, 9}
	expected := []int{-7, -5, -1, 0, 0, 2, 3, 3, 4, 6, 9, 10, 12, 20, 35}
	result := MergeSortBasic(unsorted)

	for idx, val := range expected {
		if val != result[idx] {
			t.Errorf("MergeSortBasic -- Expected %d at idx %d; got %d\n", val, idx, result[idx])
		}
	}
}

func TestMergeSortGoroutine(t *testing.T) {
	unsorted := []int{10, -5, 6, 2, -7, 20, 3, 4, -1, 0, 0, 3, 12, 35, 9}
	expected := []int{-7, -5, -1, 0, 0, 2, 3, 3, 4, 6, 9, 10, 12, 20, 35}
	result := MergeSortGoroutine(unsorted)

	for idx, val := range expected {
		if val != result[idx] {
			t.Errorf("MergeSortGoroutine -- Expected %d at idx %d; got %d\n", val, idx, result[idx])
		}
	}
}

func TestMergeSortIterative(t *testing.T) {
	unsorted := []int{10, -5, 6, 2, -7, 20, 3, 4, -1, 0, 0, 3, 12, 92, 35, -12, 9}
	expected := []int{-12, -7, -5, -1, 0, 0, 2, 3, 3, 4, 6, 9, 10, 12, 20, 35, 92}
	result := MergeSortIterative(unsorted)

	for idx, val := range expected {
		if val != result[idx] {
			t.Errorf("MergeSortIterative -- Expected %d at idx %d; got %d\n", val, idx, result[idx])
		}
	}
}

var sampleLargeArr = []int{}

func generateArray() {
	rand.Seed(42)
	if len(sampleLargeArr) > 1 {
		return
	}
	for i := 0; i < 38888; i++ {
		sampleLargeArr = append(sampleLargeArr, rand.Int())
	}
}

func clearArray() {
	sampleLargeArr = []int{}
}

func BenchmarkQuickSortBasic(b *testing.B) {
	generateArray()
	for n := 0; n < b.N; n++ {
		QuickSortBasic(sampleLargeArr)
	}
}
func BenchmarkHeapSortBasic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		HeapSortBasic(sampleLargeArr)
	}
}
func BenchmarkMergeSortBasic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortBasic(sampleLargeArr)
	}
}

func BenchmarkMergeSortGoroutine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortGoroutine(sampleLargeArr)
	}
}

func BenchmarkMergeSortIterative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortIterative(sampleLargeArr)
	}
}
