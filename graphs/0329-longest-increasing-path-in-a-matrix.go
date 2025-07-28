package graphs

// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

// Approach: Depth-First Search
// Time: O(MxN)
// Space: O(MxN)
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dirs := [5]int{-1, 0, 1, 0, -1}

	dp := make([][]int, m)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}

	var r, c int

	var dfs func(row int, col int) int

	dfs = func(row int, col int) int {
		if dp[row][col] != 0 {
			return dp[row][col]
		}

		length := 0

		for d := 1; d < 5; d++ {
			r, c = row+dirs[d-1], col+dirs[d]
			if min(r, c) >= 0 && r < m && c < n && matrix[row][col] < matrix[r][c] {
				length = max(length, dfs(r, c))
			}
		}

		length++
		dp[row][col] = length

		return length
	}
	length := 1

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			length = max(length, dfs(row, col))
		}
	}

	return length
}
