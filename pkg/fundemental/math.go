package fundemental

// Add two integers together
func Add(x int, y int) int {
	return x + y
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
