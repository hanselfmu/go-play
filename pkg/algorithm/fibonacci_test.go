package algorithm

import "testing"

func BenchmarkFib40(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(40)
	}
}

func BenchmarkFibGo40(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibGo(40)
	}
}
