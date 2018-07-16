package main

import (
	"fmt"

	a "github.com/hanselfmu/go-play/pkg/algorithm"
	f "github.com/hanselfmu/go-play/pkg/fundamental"
)

func main() {
	runfundamentals(f.Set{
		// "pointer": f.SetValStub,
		// "math":    f.SetValStub,
	})
	runAlgorithms(f.Set{
		// "fibonacci":   f.SetValStub,
		// "binary tree":         f.SetValStub,
		// "kth of sorted array": f.SetValStub,
		// "lemonade change":     f.SetValStub,
		// "jump game": f.SetValStub,
		"zig zag": f.SetValStub,
	})
}

func runfundamentals(tasks f.Set) {
	if tasks.Contains("pointer") {
		f.TestPointers()
	}

	if tasks.Contains("multireturn") {
		sortedInts, lens := f.SortInts([]int{1, 8, 4, 2, -3, 6})
		fmt.Printf("SortInts: %v\tlens: %d\n", sortedInts, lens)
	}

	if tasks.Contains("math") {
		fmt.Printf("Add(3, 4): %d\n", f.Add(3, 4))
		fmt.Printf("Sqrt(64): %f\n", f.Sqrt(64))
	}
}

func runAlgorithms(tasks f.Set) {
	treeValues := []f.Comparable{
		f.ComparableInt(20),
		f.ComparableInt(1),
		f.ComparableInt(2),
		f.ComparableInt(4),
		f.ComparableInt(-1),
		f.ComparableInt(5),
		f.ComparableInt(18),
		f.ComparableInt(-6),
		f.ComparableInt(35),
		f.ComparableInt(23),
		f.ComparableInt(12),
		f.ComparableInt(16),
		f.ComparableInt(9),
		f.ComparableInt(-12),
		f.ComparableInt(17),
		f.ComparableInt(13),
		f.ComparableInt(-16),
		f.ComparableInt(42),
		f.ComparableInt(3)}

	if tasks.Contains("fibonacci") {
		fmt.Printf("normal recursive fibonacci: %d\n", a.FibRecursiveNormal(15))
		fmt.Printf("normal iterative fibonacci: %d\n", a.FibIterativeNormal(15))
		fmt.Printf("goroutine recursive fibonacci: %d\n", a.FibRecursiveGo(15))
	}

	if tasks.Contains("binary tree") {
		bt := f.NewBinaryTree(treeValues)

		fmt.Printf("binary search tree: %v\n", bt)
		fmt.Printf("binary search tree right: %v\n", bt.Right)
		fmt.Printf("binary search tree depth recursive: %v\n", bt.DepthRec())
		fmt.Printf("binary search tree depth iterative: %v\n", bt.DepthIte())
	}

	if tasks.Contains("kth of sorted array") {
		arr1 := []int{11, 13, 16, 19, 20, 26}
		arr2 := []int{-5, -2, 1, 3, 6, 18, 21, 25}

		fmt.Printf("1st: %d\n", a.KthOfSortedArrays(arr1, arr2, 0))
		fmt.Printf("2nd: %d\n", a.KthOfSortedArrays(arr1, arr2, 1))
		fmt.Printf("3rd: %d\n", a.KthOfSortedArrays(arr1, arr2, 2))
		fmt.Printf("4th: %d\n", a.KthOfSortedArrays(arr1, arr2, 3))
		fmt.Printf("5th: %d\n", a.KthOfSortedArrays(arr1, arr2, 4))
		fmt.Printf("6th: %d\n", a.KthOfSortedArrays(arr1, arr2, 5))
		fmt.Printf("7th: %d\n", a.KthOfSortedArrays(arr1, arr2, 6))
		fmt.Printf("8th: %d\n", a.KthOfSortedArrays(arr1, arr2, 7))
		fmt.Printf("9th: %d\n", a.KthOfSortedArrays(arr1, arr2, 8))
		fmt.Printf("10th: %d\n", a.KthOfSortedArrays(arr1, arr2, 9))
		fmt.Printf("11th: %d\n", a.KthOfSortedArrays(arr1, arr2, 10))
		fmt.Printf("12th: %d\n", a.KthOfSortedArrays(arr1, arr2, 11))
		fmt.Printf("13th: %d\n", a.KthOfSortedArrays(arr1, arr2, 12))
		fmt.Printf("14th: %d\n", a.KthOfSortedArrays(arr1, arr2, 13))
	}

	if tasks.Contains("kth of sorted array") {
		arr1 := []int{5, 5, 5, 10, 20}
		arr2 := []int{5, 10, 5, 5, 20}

		fmt.Printf("can complete seq 1? %v\n", a.LemonadeChange(arr1))
		fmt.Printf("can complete seq 2? %v\n", a.LemonadeChange(arr2))
	}

	if tasks.Contains("jump game") {
		nums1 := []int{2, 3, 1, 1, 4}
		fmt.Printf("can jump %v\t?: %v\n", nums1, a.CanJump(nums1))
	}

	if tasks.Contains("zig zag") {
		s := "ABCDE"
		fmt.Printf("%s with rows: %d: %s\n", s, 3, a.ConvertZigZag(s, 4))
	}
}
