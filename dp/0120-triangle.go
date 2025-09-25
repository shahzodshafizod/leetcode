package dp

// https://leetcode.com/problems/triangle/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	dp := make([]int, n+1)
	for row := n - 1; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			dp[col] = triangle[row][col] + min(dp[col], dp[col+1])
		}
	}

	return dp[0]
}
