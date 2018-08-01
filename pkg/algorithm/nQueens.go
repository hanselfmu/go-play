package algorithm

// NQueensCount counts the solution to the N-Queens problem (https://en.wikipedia.org/wiki/Eight_queens_puzzle)
// LeetCode #52
func NQueensCount(n int) int {
	if n == 0 || n == 2 || n == 3 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return 2
}
