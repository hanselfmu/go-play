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
		"binary tree":         f.SetValStub,
		"kth of sorted array": f.SetValStub,
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
	fmt.Println("testing this shit", 1/2)
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

		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 0))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 1))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 2))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 3))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 4))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 5))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 6))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 7))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 8))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 9))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 10))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 11))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 12))
		fmt.Printf("kth: %d\n", a.KthOfSortedArrays(arr1, arr2, 13))
	}
}
