package main

import (
	"fmt"

	"github.com/hanselfmu/go-play/pkg/algorithm"
	"github.com/hanselfmu/go-play/pkg/fundemental"
)

func main() {
	runFundementals()
	runAlgorithms()
}

func runFundementals() {
	fmt.Println("running fundementals")
	fmt.Printf("result from Add(3 + 4): %d\n", fundemental.Add(3, 4))
	fmt.Println(fundemental.SortInts([]int{1, 8, 4, 2, -3, 6}))
}

func runAlgorithms() {
	fmt.Println("running algorithms")
	fmt.Println(algorithm.Fibonacci(3))
}
