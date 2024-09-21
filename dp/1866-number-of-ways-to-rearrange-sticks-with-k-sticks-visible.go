package dp

// https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/

// Approach: Bottom-Up Dynamic Programming
// Time: O(NxK)
// Space: O(K)
func rearrangeSticks(n int, k int) int {
	const mod int = 1e9 + 7
	var dp = make([]int, k+1)
	dp[0] = 1
	for x := 1; x <= n; x++ {
		tmp := make([]int, k+1)
		for y := 1; y <= k; y++ {
			tmp[y] = (dp[y-1] + (x-1)*dp[y]) % mod
		}
		dp = tmp
	}
	return dp[k]
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(NxK)
// // Space: O(NxK)
// func rearrangeSticks(n int, k int) int {
// 	const mod int = 1e9 + 7
// 	var dp = make([][]int, n+1)
// 	for idx := range dp {
// 		dp[idx] = make([]int, k+1)
// 	}
// 	var dfs func(n int, k int) int
// 	dfs = func(n int, k int) int {
// 		if n == k {
// 			return 1
// 		}
// 		if n == 0 || k == 0 {
// 			return 0
// 		}
// 		if dp[n][k] != 0 {
// 			return dp[n][k]
// 		}
// 		// 1. decision to include
// 		// if we put largest (n) in the end,
// 		// it'll be visible, so we decrease k
// 		count := dfs(n-1, k-1)
// 		// 2. decision NOT to include
// 		// non-largest will be hidden by the largest
// 		count += (n - 1) * dfs(n-1, k)
// 		count %= mod
// 		dp[n][k] = count
// 		return count
// 	}
// 	return dfs(n, k)
// }
