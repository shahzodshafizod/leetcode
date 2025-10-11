package dp

// https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/

// Approach #2: Bottom-Up Dynamic Programming + Prefix Sum
// Time: O(mnk*(m+n)), m=len(pizza), n=len(pizza[0])
// Space: O(mnk)
func ways(pizza []string, k int) int {
	const mod = int(1e9) + 7

	rows, cols := len(pizza), len(pizza[0])
	presum := make([][]int, rows+1)
	dp := make([][][]int, rows+1)

	for r := range rows + 1 {
		presum[r] = make([]int, cols+1)
		dp[r] = make([][]int, cols+1)

		for c := range cols + 1 {
			dp[r][c] = make([]int, k)
		}
	}

	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			if pizza[r][c] == 'A' {
				presum[r][c] = 1
			}

			presum[r][c] += presum[r+1][c] + presum[r][c+1] - presum[r+1][c+1]
		}
	}

	for r := range rows {
		for c := range cols {
			if presum[r][c] > 0 {
				dp[r][c][0] = 1
			}
		}
	}

	for cuts := 1; cuts < k; cuts++ {
		for r := range rows {
			for c := range cols {
				for nr := r + 1; nr < rows; nr++ {
					if presum[r][c] > presum[nr][c] && presum[nr][c] >= cuts {
						dp[r][c][cuts] = (dp[r][c][cuts] + dp[nr][c][cuts-1]) % mod
					}
				}

				for nc := c + 1; nc < cols; nc++ {
					if presum[r][c] > presum[r][nc] && presum[r][nc] >= cuts {
						dp[r][c][cuts] = (dp[r][c][cuts] + dp[r][nc][cuts-1]) % mod
					}
				}
			}
		}
	}

	return dp[0][0][k-1]
}

// // Approach #1: Top-Down Dynamic Programming + Prefix Sum
// // Time: O(mnk*(m+n)), m=len(pizza), n=len(pizza[0])
// // Space: O(mnk)
// func ways(pizza []string, k int) int {
// 	const mod = int(1e9) + 7

// 	rows, cols := len(pizza), len(pizza[0])
// 	presum := make([][]int, rows+1)

// 	memo := make([][][]*int, rows+1)
// 	for row := range rows + 1 {
// 		presum[row] = make([]int, cols+1)

// 		memo[row] = make([][]*int, cols+1)
// 		for col := range cols + 1 {
// 			memo[row][col] = make([]*int, k)
// 		}
// 	}

// 	for row := rows - 1; row >= 0; row-- {
// 		for col := cols - 1; col >= 0; col-- {
// 			if pizza[row][col] == 'A' {
// 				presum[row][col] = 1
// 			}

// 			presum[row][col] += presum[row+1][col] + presum[row][col+1] - presum[row+1][col+1]
// 		}
// 	}

// 	var dp func(row int, col int, k int) int

// 	dp = func(row int, col int, k int) int {
// 		if memo[row][col][k] != nil {
// 			return *memo[row][col][k]
// 		}

// 		if presum[row][col] == 0 {
// 			return 0
// 		}

// 		if k == 0 {
// 			return 1
// 		}

// 		var cnt int
// 		// cutting horizontally
// 		for nrow := row + 1; nrow < rows; nrow++ {
// 			if presum[row][col] > presum[nrow][col] {
// 				cnt = (cnt + dp(nrow, col, k-1)) % mod
// 			}
// 		}
// 		// cutting vertically
// 		for ncol := col + 1; ncol < cols; ncol++ {
// 			if presum[row][col] > presum[row][ncol] {
// 				cnt = (cnt + dp(row, ncol, k-1)) % mod
// 			}
// 		}

// 		memo[row][col][k] = &cnt

// 		return cnt
// 	}

// 	return dp(0, 0, k-1)
// }
