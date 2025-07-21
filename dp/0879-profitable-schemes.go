package dp

// https://leetcode.com/problems/profitable-schemes/

// Approach: Bottom-Up Dynamic Programming
// Time: O(NGP), G=len(group), P=minProfit
// Space: O(NP)
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	const mod int = 1e9 + 7
	// dp[i][j] means the count of schemes with i profit and j members.
	dp := make([][]int, minProfit+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}
	dp[0][0] = 1
	g := len(group)
	for k := 0; k < g; k++ { // O(G)
		g := group[k]
		p := profit[k]
		for i := minProfit; i >= 0; i-- { // O(P)
			for j := n - g; j >= 0; j-- { // O(N)
				dp[min(minProfit, i+p)][j+g] = (dp[min(minProfit, i+p)][j+g] + dp[i][j]) % mod
			}
		}
	}
	count := 0
	for _, cnt := range dp[minProfit] {
		count = (count + cnt) % mod
	}
	return count
}

// // Approach: Top-Down Dynamic Programming (Time Limit Exceeded)
// // Time: O(NGP), G=len(group), P=minProfit
// // Space: O(NGP)
// func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
// 	var m = len(profit)
// 	const mod int = 1e9 + 7
// 	var memo [101][101][101]int
// 	var dfs func(idx int, members int, profits int) int
// 	dfs = func(idx int, members int, profits int) int {
// 		if idx == m {
// 			if profits == minProfit {
// 				return 1
// 			}
// 			return 0
// 		}
// 		if memo[idx][members][profits] != 0 {
// 			return memo[idx][members][profits]
// 		}
// 		// decision NOT to include
// 		var count = dfs(idx+1, members, profits)
// 		// decision to include
// 		if members+group[idx] <= n {
// 			count = (count + dfs(
// 				idx+1,
// 				members+group[idx],
// 				min(minProfit, profits+profit[idx]),
// 			)) % mod
// 		}
// 		memo[idx][members][profits] = count
// 		return count
// 	}
// 	return dfs(0, 0, 0)
// }
