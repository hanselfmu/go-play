package algorithm

// MostWaterContainer calculates the max area according to the given heights array.
// LeetCode #11
func MostWaterContainer(height []int) int {
	n := len(height)

	max := 0
	for li, ri := 0, n-1; li < ri; {
		distance := ri - li
		// Whichever is the smaller is having the max possible value it can have
		// given that the other side cannot move farther away
		if height[li] < height[ri] {
			cur := height[li] * distance
			if cur > max {
				max = cur
			}
			li++
		} else {
			cur := height[ri] * distance
			if cur > max {
				max = cur
			}
			ri--
		}
	}

	return max
}
