package dp

// https://leetcode.com/problems/cherry-pickup/

// Hint: Run two robots from start to end. Simultaneous DFS.

// Approach #3: Dynamic Programming (Bottom-Up)
// Time: O(N^3)
// Space: O(N^2)
func cherryPickup(grid [][]int) int {
	n := len(grid)
	m := (n << 1) - 1 // 2*n-1: # of steps to reach (n-1,n-1)

	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for t := 1; t < m; t++ {
		for i := n - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				if t-i < 0 || t-i >= n || t-j < 0 || t-j >= n ||
					grid[i][t-i] == -1 || grid[j][t-j] == -1 {
					dp[i][j] = -1

					continue
				}

				if i > 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j])
				}

				if j > 0 {
					dp[i][j] = max(dp[i][j], dp[i][j-1])
				}

				if i > 0 && j > 0 {
					dp[i][j] = max(dp[i][j], dp[i-1][j-1])
				}

				if dp[i][j] >= 0 {
					dp[i][j] += grid[i][t-i]
					if i != j {
						dp[i][j] += grid[j][t-j]
					}
				}
			}
		}
	}

	return max(0, dp[n-1][n-1])
}

// // Approach #2: Dynamic Programming (Top-Down)
// // Time: O(N^3)
// // Space: O(N^3)
// func cherryPickup(grid [][]int) int {
// 	// (r1+c1) = (r2+c2)
// 	// c2 = r1+c1-r2
// 	var n = len(grid)
// 	var memo [51][51][51]*int
// 	var dfs func(r1 int, c1 int, r2 int) int
// 	dfs = func(r1 int, c1 int, r2 int) int {
// 		var c2 = r1 + c1 - r2
// 		if r1 == n || c1 == n || r2 == n || c2 == n ||
// 			grid[r1][c1] == -1 || grid[r2][c2] == -1 {
// 			// this -ve should make result of other instances equal to <= 0
// 			return -300 // or -3*(50+49)+1
// 		}
// 		if r1 == n-1 && c1 == n-1 {
// 			// (r1,c1) & (r2,c2)
// 			// both come to the end
// 			// at the same time
// 			return grid[r2][c2]
// 		}
// 		if memo[r1][c1][r2] != nil {
// 			return *memo[r1][c1][r2]
// 		}
// 		var count = grid[r1][c1]
// 		if c1 != c2 {
// 			count += grid[r2][c2]
// 		}
// 		// person1 -> down, person2 -> down
// 		downDown := dfs(r1+1, c1, r2+1) // ,c2
// 		// person1 -> down, person2 -> right
// 		downRight := dfs(r1+1, c1, r2) // c2+1
// 		// person1 -> right, person2 -> down
// 		rightDown := dfs(r1, c1+1, r2+1) // c2
// 		// person1 -> right, person2 -> right
// 		rightRight := dfs(r1, c1+1, r2) // c2+1
// 		count += max(downDown, downRight, rightDown, rightRight)
// 		memo[r1][c1][r2] = &count
// 		return count
// 	}
// 	return max(0, dfs(0, 0, 0))
// }

// // Approach #1: Dynamic Programming (Top-Down)
// // Time: O(N^4)
// // Space: O(N^4)
// func cherryPickup(grid [][]int) int {
// 	var n = len(grid)
// 	var memo [51][51][51][51]*int
// 	var dfs func(r1 int, c1 int, r2 int, c2 int) int
// 	dfs = func(r1 int, c1 int, r2 int, c2 int) int {
// 		if r1 == n || c1 == n || r2 == n || c2 == n ||
// 			grid[r1][c1] == -1 || grid[r2][c2] == -1 {
// 			// this -ve should make result of other instances equal to <= 0
// 			return -300 // or -3*(50+49)+1
// 		}
// 		if r1 == n-1 && c1 == n-1 {
// 			// (r1,c1) & (r2,c2)
// 			// both come to the end
// 			// at the same time
// 			return grid[r2][c2]
// 		}
// 		if memo[r1][c1][r2][c2] != nil {
// 			return *memo[r1][c1][r2][c2]
// 		}
// 		var count = grid[r1][c1]
// 		if c1 != c2 {
// 			count += grid[r2][c2]
// 		}
// 		// person1 -> down, person2 -> down
// 		downDown := dfs(r1+1, c1, r2+1, c2)
// 		// person1 -> down, person2 -> right
// 		downRight := dfs(r1+1, c1, r2, c2+1)
// 		// person1 -> right, person2 -> down
// 		rightDown := dfs(r1, c1+1, r2+1, c2)
// 		// person1 -> right, person2 -> right
// 		rightRight := dfs(r1, c1+1, r2, c2+1)
// 		count += max(downDown, downRight, rightDown, rightRight)
// 		memo[r1][c1][r2][c2] = &count
// 		return count
// 	}
// 	return max(0, dfs(0, 0, 0, 0))
// }
