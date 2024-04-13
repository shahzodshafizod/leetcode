package dp

// https://leetcode.com/problems/maximal-square/

// time: O(rows*cols)
// space: O(rows*cols)
func maximalSquare(matrix [][]byte) int {
	var rows, cols = len(matrix), len(matrix[0])
	var dp = make([][]int, rows+1)
	for idx := range dp {
		dp[idx] = make([]int, cols+1)
	}
	var maxside = 0
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if matrix[row][col] == '1' {
				dp[row][col] = 1 + min(dp[row+1][col+1],
					min(dp[row][col+1], dp[row+1][col]))
				maxside = max(maxside, dp[row][col])
			}
		}
	}
	return maxside * maxside
}

// // time: O(rows*cols)
// // space: O(rows*cols)
// func maximalSquare(matrix [][]byte) int {
// 	var rows, cols = len(matrix), len(matrix[0])
// 	var cache = make([][]*int, rows)
// 	for idx := range cache {
// 		cache[idx] = make([]*int, cols)
// 	}
// 	var dp func(row int, col int) int
// 	dp = func(row int, col int) int {
// 		if row == rows || col == cols || matrix[row][col] == '0' {
// 			return 0
// 		}
// 		if cache[row][col] != nil {
// 			return *cache[row][col]
// 		}
// 		var side = 1 + min(dp(row+1, col+1),
// 			min(dp(row+1, col), dp(row, col+1)))
// 		cache[row][col] = &side
// 		return side
// 	}
// 	var maxside = 0
// 	for row := 0; row < rows; row++ {
// 		for col := 0; col < cols; col++ {
// 			if matrix[row][col] == '1' {
// 				maxside = max(maxside, dp(row, col))
// 			}
// 		}
// 	}
// 	return maxside * maxside
// }
