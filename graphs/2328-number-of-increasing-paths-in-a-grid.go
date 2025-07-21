package graphs

// https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/

// Approach: Depth-First Search
// Time: O(MxN)
// Space: O(MxN)
func countPaths(grid [][]int) int {
	const mod = int(1e9 + 7)
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}
	dirs := [5]int{-1, 0, 1, 0, -1}
	var r, c int
	var dfs func(row int, col int) int
	dfs = func(row int, col int) int {
		if dp[row][col] != 0 {
			return dp[row][col]
		}
		count := 1
		for d := 1; d < 5; d++ {
			r, c = row+dirs[d-1], col+dirs[d]
			if min(r, c) >= 0 && r < m && c < n &&
				grid[row][col] < grid[r][c] {
				count = (count + dfs(r, c)) % mod
			}
		}
		dp[row][col] = count
		return count
	}
	count := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			count = (count + dfs(row, col)) % mod
		}
	}
	return count
}
