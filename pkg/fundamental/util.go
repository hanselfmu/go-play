package fundamental

import (
	"sort"
)

// SortInts proxies sort pkg's Ints sort
func SortInts(ints []int) ([]int, int) {
	sort.Ints(ints)
	return ints, len(ints)
}
