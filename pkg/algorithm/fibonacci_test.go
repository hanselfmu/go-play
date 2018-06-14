package algorithm

import "testing"

func BenchmarkFibRecNormal15(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibRecursiveNormal(15)
	}
}

func BenchmarkFibIteNormal15(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibIterativeNormal(15)
	}
}

func BenchmarkFibRecGo15(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibRecursiveGo(15)
	}
}
