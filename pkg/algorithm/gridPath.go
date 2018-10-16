package algorithm

import (
	"math"
)

// UniquePaths is the easiest and the most basic grid problem of all.
// LeetCode #62
func UniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	m--
	n--

	if m > n {
		return factorialWithRange(m+1, m+n) / factorialWithRange(1, n)
	} else {
		return factorialWithRange(n+1, m+n) / factorialWithRange(1, m)
	}
}

func factorialWithRange(start, end int) int {
	res := 1
	for v := start; v <= end; v++ {
		res *= v
	}
	return res
}

// MinGridPathSum returns the sum of the minimum path from top left to bottom right
// for a grid.
// LeetCode #64
func MinGridPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	colSize := len(grid[0])
	prev := make([]int, colSize)
	// initialize
	for colIdx, val := range grid[0] {
		if colIdx == 0 {
			prev[0] = val
		} else {
			prev[colIdx] = prev[colIdx-1] + val
		}
	}

	for _, row := range grid[1:] {
		cur := make([]int, colSize)
		for colIdx, val := range row {
			if colIdx == 0 {
				cur[0] = prev[0] + val
			} else {
				if cur[colIdx-1] < prev[colIdx] {
					cur[colIdx] = cur[colIdx-1] + val
				} else {
					cur[colIdx] = prev[colIdx] + val
				}
			}
		}
		prev = cur
	}

	return prev[colSize-1]
}

// DungeonGame calculates the minimum health required for a knight to pass a dungeon
// from top left to bottom right in a grid.
// LeetCode #174
func DungeonGame(dungeon [][]int) int {
	rowSize := len(dungeon)
	if rowSize == 0 {
		return 0
	}
	colSize := len(dungeon[0])
	prevRow := make([]int, colSize)

	start := 0
	if dungeon[rowSize-1][colSize-1] < 0 {
		start = -1 * dungeon[rowSize-1][colSize-1]
	}

	for row := rowSize - 1; row >= 0; row-- {
		curRow := make([]int, colSize)
		for col := colSize - 1; col >= 0; col-- {
			if col == colSize-1 && row == rowSize-1 {
				curRow[col] = start
				continue
			}
			cell := dungeon[row][col]

			left, up := math.MinInt32, math.MinInt32
			if row < rowSize-1 {
				left = cell - prevRow[col]
			}
			if col < colSize-1 {
				up = cell - curRow[col+1]
			}
			if left > up {
				curRow[col] = left
			} else {
				curRow[col] = up
			}
			if curRow[col] > 0 {
				curRow[col] = 0
			} else {
				curRow[col] *= -1
			}
		}

		prevRow = curRow
	}

	return prevRow[0] + 1
}
