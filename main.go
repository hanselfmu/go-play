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
		"binary tree": f.SetValStub,
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
	if tasks.Contains("fibonacci") {
		fmt.Printf("normal recursive fibonacci: %d\n", a.FibRecursiveNormal(15))
		fmt.Printf("normal iterative fibonacci: %d\n", a.FibIterativeNormal(15))
		fmt.Printf("goroutine recursive fibonacci: %d\n", a.FibRecursiveGo(15))
	}

	if tasks.Contains("binary tree") {
		bt := f.NewBinaryTree([]f.Comparable{
			f.ComparableInt(1),
			f.ComparableInt(2),
			f.ComparableInt(4),
			f.ComparableInt(-1),
			f.ComparableInt(5),
			f.ComparableInt(3)})

		fmt.Printf("binary search tree: %v\n", bt)
		fmt.Printf("binary search tree right: %v\n", bt.Right)
		fmt.Printf("binary search tree depth recursive: %v\n", bt.DepthRec())
		fmt.Printf("binary search tree depth iterative: %v\n", bt.DepthIte())
	}
}
