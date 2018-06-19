package main

import (
	"github.com/hanselfmu/go-play/pkg/fundemental"
)

func main() {
	runFundementals()
	runAlgorithms()
}

func runFundementals() {
	fundemental.TestPointers()
	// fmt.Println("running fundementals")
	// fmt.Printf("Add(3 + 4): %d\n", fundemental.Add(3, 4))
	// sortedInts, lens := fundemental.SortInts([]int{1, 8, 4, 2, -3, 6})
	// fmt.Printf("SortInts: %v\tlens: %d\n", sortedInts, lens)
	// fmt.Printf("Sqrt(64): %f\n", fundemental.Sqrt(64))
}

func runAlgorithms() {
	// fmt.Println("running algorithms")
	// fmt.Printf("normal recursive fibonacci: %d\n", algorithm.FibRecursiveNormal(15))
	// fmt.Printf("normal iterative fibonacci: %d\n", algorithm.FibIterativeNormal(15))
	// fmt.Printf("goroutine recursive fibonacci: %d\n", algorithm.FibRecursiveGo(15))
}
