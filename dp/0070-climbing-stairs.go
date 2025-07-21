package dp

// https://leetcode.com/problems/climbing-stairs/

// Bottom-Up
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// // Top-Down
// func climbStairs(n int) int {
// 	return climbStairsMemoization(n, make([]int, n))
// }

// func climbStairsMemoization(n int, dp []int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	if dp[n-1] != 0 {
// 		return dp[n-1]
// 	}
// 	dp[n-1] = climbStairsMemoization(n-1, dp) + climbStairsMemoization(n-2, dp)
// 	return dp[n-1]
// }
