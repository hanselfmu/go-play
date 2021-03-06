package fundamental

import (
	"reflect"
)

// IntSize records size of integer on the current architecture in bytes
var IntSize = reflect.TypeOf(0).Size()

// Add two integers together
func Add(x int, y int) int {
	return x + y
}

// Abs calculates the absolute number of x
func Abs(x int) int {
	mask := x >> (IntSize*8 - 1)
	return (x + mask) ^ mask
}

// Max returns the bigger number between x and y
func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// Sqrt computes square root for a float
func Sqrt(x float64) float64 {
	const MinChange = 0.0001
	z := 1.0
	diff := 100.0
	for diff > MinChange || diff < -MinChange {
		diff = (z*z - x) / (2 * z)

		z -= diff
	}

	return z
}
