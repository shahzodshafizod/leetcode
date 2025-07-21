package dp

// https://leetcode.com/problems/fibonacci-number/

// Bottom-Up
func fib(n int) int {
	if n <= 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// // Top-Down
// func fib(n int) int {
// 	return fibMemoization(n, make([]int, n))
// }

// func fibMemoization(n int, dp []int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	if dp[n-1] != 0 {
// 		return dp[n-1]
// 	}
// 	dp[n-1] = fibMemoization(n-1, dp) + fibMemoization(n-2, dp)
// 	return dp[n-1]
// }
