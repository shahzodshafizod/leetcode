package bits

/*
There are only two rules that you need to know in this problem:
	If the first number in the row is 0, flip the row.
	If the count of 0 in the col is greater than the count of 1, flip the col.
*/

// https://leetcode.com/problems/score-after-flipping-matrix/

// time: O(m x n)
// space: O(1)
func matrixScore(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	for row := 0; row < m; row++ { // O(m x n)
		if grid[row][0] == 0 {
			for col := 0; col < n; col++ {
				grid[row][col] ^= 1
			}
		}
	}
	var score = 0
	var zeros, ones int
	var bit = 1 << (n - 1)
	for col := 0; col < n; col++ { // O(n x m)
		zeros = 0
		for row := 0; row < m; row++ {
			if grid[row][col] == 0 {
				zeros++
			}
		}
		ones = m - zeros
		score += max(ones, zeros) * bit
		bit >>= 1
	}
	return score
}
