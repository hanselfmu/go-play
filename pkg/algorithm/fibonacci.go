package algorithm

import (
	"runtime"
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

// FibIterativeNormal calculates finobacci for input n with iteration
func FibIterativeNormal(n int) int {
	if n <= 1 {
		return n
	}

	x, y := 0, 1
	for i := 2; i <= n; i++ {
		y = x + y
		x = y - x
	}

	return y
}

func fibRecursiveGo(n int, c chan int, wg *sync.WaitGroup) {

	wg.Add(1)
	defer wg.Done()

	if n <= 1 {
		c <- n
	} else {
		go fibRecursiveGo(n-1, c, wg)
		go fibRecursiveGo(n-2, c, wg)
		r1, r2 := <-c, <-c
		c <- r1 + r2
	}
}

// FibRecursiveGo uses goroutines to speed up fibonacci calculation
func FibRecursiveGo(n int) int {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	c := make(chan int, 1) // here is tricky; a bufferless channel won't work
	fibRecursiveGo(n, c, &wg)

	wg.Wait()
	r := <-c
	return r
}
