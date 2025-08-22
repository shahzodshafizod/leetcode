package dp

// https://leetcode.com/problems/count-square-submatrices-with-all-ones/

// Approach #4: Bottom-Up Dynamic Programming (in-place)
// Time: O(mn)
// Space: O(1)
func countSquares(matrix [][]int) int {
	var (
		m, n  = len(matrix), len(matrix[0])
		count = 0
	)

	for row := range m {
		for col := range n {
			if row > 0 && col > 0 && matrix[row][col] == 1 {
				matrix[row][col] = 1 + min(
					matrix[row-1][col],
					matrix[row][col-1],
					matrix[row-1][col-1],
				)
			}

			count += matrix[row][col]
		}
	}

	return count
}

// // Approach #3: Bottom-Up Dynamic Programming (Optimized Space)
// // Time: O(mn)
// // Space: O(n)
// func countSquares(matrix [][]int) int {
// 	var m, n = len(matrix), len(matrix[0])
// 	var prev = make([]int, n+1)
// 	var curr = make([]int, n+1)
// 	var count = 0
// 	for row := 1; row <= m; row++ {
// 		prev, curr = curr, prev
// 		for col := 1; col <= n; col++ {
// 			curr[col] = matrix[row-1][col-1]
// 			if matrix[row-1][col-1] == 1 {
// 				curr[col] += min(
// 					prev[col],
// 					curr[col-1],
// 					prev[col-1],
// 				)
// 			}
// 			count += curr[col]
// 		}
// 	}
// 	return count
// }

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(mn)
// // Space: O(mn)
// func countSquares(matrix [][]int) int {
// 	var m, n = len(matrix), len(matrix[0])
// 	var dp = make([][]int, m+1)
// 	for row := range dp {
// 		dp[row] = make([]int, n+1)
// 	}
// 	var count = 0
// 	for row := 1; row <= m; row++ {
// 		for col := 1; col <= n; col++ {
// 			if matrix[row-1][col-1] == 1 {
// 				dp[row][col] = 1 + min(
// 					dp[row-1][col],
// 					dp[row][col-1],
// 					dp[row-1][col-1],
// 				)
// 			}
// 			count += dp[row][col]
// 		}
// 	}
// 	return count
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(mn)
// // Space: O(mn)
// func countSquares(matrix [][]int) int {
// 	var m, n = len(matrix), len(matrix[0])
// 	var cache = make([][]*int, m)
// 	for row := range cache {
// 		cache[row] = make([]*int, n)
// 	}
// 	var dp func(row int, col int) int
// 	dp = func(row int, col int) int {
// 		if row == m || col == n || matrix[row][col] == 0 {
// 			return 0
// 		}
// 		if cache[row][col] != nil {
// 			return *cache[row][col]
// 		}
// 		cache[row][col] = new(int)
// 		*cache[row][col] = 1 + min(
// 			dp(row, col+1),
// 			dp(row+1, col),
// 			dp(row+1, col+1),
// 		)
// 		return *cache[row][col]
// 	}
// 	var count = 0
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			count += dp(row, col)
// 		}
// 	}
// 	return count
// }
