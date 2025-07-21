package dp

// https://leetcode.com/problems/dungeon-game/

// Approach: Dynamic Programming (Bottom-Up)
// Time: O(MxN)
// Space: O(MxN)
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for row := range dp {
		dp[row] = make([]int, n+1)
		for col := range dp[row] {
			dp[row][col] = 1e9
		}
	}
	dp[m-1][n] = 1
	dp[m][n-1] = 1
	for row := m - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			dp[row][col] = max(1,
				min(
					dp[row][col+1],
					dp[row+1][col],
				)-dungeon[row][col],
			)
		}
	}
	return dp[0][0]
}

// // Approach: Dynamic Programming (Top-Down)
// // Time: O(MxN)
// // Space: O(MxN)
// func calculateMinimumHP(dungeon [][]int) int {
// 	var m, n = len(dungeon), len(dungeon[0])
// 	var dp = make([][]int, m)
// 	for idx := range dp {
// 		dp[idx] = make([]int, n)
// 	}
// 	var dfs func(row int, col int) int
// 	dfs = func(row int, col int) int {
// 		if row == m || col == n {
// 			return 1e9
// 		}
// 		if row == m-1 && col == n-1 {
// 			return max(1, 1-dungeon[row][col])
// 		}
// 		if dp[row][col] != 0 {
// 			return dp[row][col]
// 		}
// 		var right = dfs(row, col+1)
// 		var down = dfs(row+1, col)
// 		dp[row][col] = max(1, min(right, down)-dungeon[row][col])
// 		return dp[row][col]
// 	}
// 	return dfs(0, 0)
// }
