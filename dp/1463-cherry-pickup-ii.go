package dp

// https://leetcode.com/problems/cherry-pickup-ii/

// Approach: Dynamic Programming (Bottom-Up)
// Time: O(m*n^2)
// Space: O(n^2)
func cherryPickupII(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var dp [2][71][71]int
	var curr, next = 0, 1
	var count int
	for row := m - 1; row >= 0; row-- {
		next, curr = curr, next
		for col1 := 0; col1 < n-1; col1++ {
			for col2 := col1 + 1; col2 < n; col2++ {
				count = 0
				for nc1 := max(0, col1-1); nc1 <= min(n-1, col1+1); nc1++ {
					for nc2 := max(0, col2-1); nc2 <= min(n-1, col2+1); nc2++ {
						count = max(count, dp[next][nc1][nc2])
					}
				}
				dp[curr][col1][col2] = count + grid[row][col1] + grid[row][col2]
			}
		}
	}
	return dp[curr][0][n-1]
}

// // Approach: Dynamic Programming (Top-Down)
// // Time: O(m*n^2)
// // Space: O(m*n^2)
// func cherryPickupII(grid [][]int) int {
// 	var m, n = len(grid), len(grid[0])
// 	var memo [71][71][71]*int
// 	var dfs func(row int, col1 int, col2 int) int
// 	dfs = func(row int, col1 int, col2 int) int {
// 		// row1 & row2 are equal
// 		if col1 == col2 || min(col1, col2) < 0 || max(col1, col2) == n {
// 			return 0
// 		}
// 		if memo[row][col1][col2] != nil {
// 			return *memo[row][col1][col2]
// 		}
// 		if row == m-1 {
// 			return grid[row][col1] + grid[row][col2]
// 		}
// 		var count = 0
// 		for nextCol1 := col1 - 1; nextCol1 <= col1+1; nextCol1++ {
// 			for nextCol2 := col2 - 1; nextCol2 <= col2+1; nextCol2++ {
// 				count = max(count, dfs(row+1, nextCol1, nextCol2))
// 			}
// 		}
// 		count += grid[row][col1] + grid[row][col2]
// 		memo[row][col1][col2] = &count
// 		return count
// 	}
// 	return dfs(0, 0, n-1)
// }
