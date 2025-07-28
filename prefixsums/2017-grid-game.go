package prefixsums

import "math"

// https://leetcode.com/problems/grid-game/

// Approach: Prefix Sum
// Time: O(n)
// Space: O(1)
func gridGame(grid [][]int) int64 {
	var top, bottom int64 = 0, 0
	for _, num := range grid[1] {
		bottom += int64(num)
	}

	var robot2 int64 = math.MaxInt64

	for col := len(grid[0]) - 1; col >= 0; col-- {
		bottom -= int64(grid[1][col])
		robot2 = min(robot2, max(top, bottom))
		top += int64(grid[0][col])
	}

	return robot2
}
