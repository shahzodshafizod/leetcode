package dp

// https://leetcode.com/problems/distinct-subsequences-ii/

// Approach: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(n)
// Space: O(n)
func distinctSubseqII(s string) int {
	const MOD int = 1e9 + 7
	dp := 1
	var last [26]int
	var key int
	for idx := range s {
		key = int(s[idx] - 'a')
		last[key], dp = dp, (dp*2-last[key])%MOD
	}
	dp-- // -1, excluding the empty subsequence
	if dp < 0 {
		dp += MOD
	}
	return dp
}

// // Approach: Bottom-Up Dynamic Programming
// // Time: O(n)
// // Space: O(n)
// func distinctSubseqII(s string) int {
// 	const MOD int = 1e9 + 7
// 	var n = len(s)
// 	var dp = make([]int, n+1)
// 	dp[0] = 1
// 	var last [26]int
// 	for idx := range last {
// 		last[idx] = n
// 	}
// 	var key int
// 	for idx := range s {
// 		key = int(s[idx] - 'a')
// 		dp[idx+1] = (dp[idx]*2 - dp[last[key]] + MOD) % MOD
// 		last[key] = idx
// 	}
// 	// -1, excluding the empty subsequence
// 	return (dp[n] - 1 + MOD) % MOD
// }

// // Approach: Top-Down Dynamic Programming
// // Time: O(n)
// // Space: O(n)
// func distinctSubseqII(s string) int {
// 	const MOD int = 1e9 + 7
// 	var last [26]int
// 	for idx := range last {
// 		last[idx] = -1
// 	}
// 	var n = len(s)
// 	var memo = make([]*int, n+1)
// 	var dp func(idx int) int
// 	dp = func(idx int) int {
// 		if idx < 0 {
// 			return 0
// 		}
// 		if idx == 0 {
// 			return 1
// 		}
// 		if memo[idx] != nil {
// 			return *memo[idx]
// 		}
// 		memo[idx] = new(int)
// 		var key = int(s[idx-1] - 'a')
// 		*memo[idx] = (dp(idx-1)*2 - dp(last[key]-1) + MOD) % MOD
// 		last[key] = idx
// 		return *memo[idx]
// 	}
//  // -1, excluding the empty subsequence
// 	return (dp(n) - 1 + MOD) % MOD
// }
