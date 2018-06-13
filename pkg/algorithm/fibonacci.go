package algorithm

import (
	"fmt"
	"sync"
)

// FibRecursiveNormal calculates finobacci for input n with recursion;
// Ideally Fibonacci sequence should be calculated with an iterative approach.
func FibRecursiveNormal(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return FibRecursiveNormal(n-1) + FibRecursiveNormal(n-2)
	}
}

func fibRecursiveGo(n int, marker string, c chan int, wg *sync.WaitGroup) {
	fmt.Printf("fibbing %d\tmarker: %s\n", n, marker)

	if n <= 1 {
		c <- n
	} else {
		go fibRecursiveGo(n-1, "1", c, wg)
		go fibRecursiveGo(n-2, "2", c, wg)
		r1 := <-c
		r2 := <-c
		fmt.Printf("r1: %d\tr2: %d\n", r1, r2)
		c <- r1 + r2
	}
}

// FibRecursiveGo uses goroutines to speed up fibonacci calculation
func FibRecursiveGo(n int) int {
	var wg sync.WaitGroup
	wg.Add(1)

	c := make(chan int, 100000)
	go fibRecursiveGo(n, "1", c, &wg)

	wg.Wait()
	fmt.Println("im here")
	r := <-c
	return r
}
