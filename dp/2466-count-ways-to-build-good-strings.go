package dp

// https://leetcode.com/problems/count-ways-to-build-good-strings/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(high)
// Space: O(high)
func countGoodStrings(low int, high int, zero int, one int) int {
	var MOD int = 1e9 + 7
	var dp = make([]int, high+1)
	dp[0] = 1
	var count = 0
	for length := 1; length <= high; length++ {
		if length-zero >= 0 {
			dp[length] = dp[length-zero]
		}
		if length-one >= 0 {
			dp[length] = (dp[length] + dp[length-one]) % MOD
		}
		if length >= low {
			count = (count + dp[length]) % MOD
		}
	}
	return count
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(high)
// // Space: O(high)
// func countGoodStrings(low int, high int, zero int, one int) int {
// 	const MOD int = 1e9 + 7
// 	var memo = make([]*int, high+1)
// 	var dp func(length int) int
// 	dp = func(length int) int {
// 		if memo[length] != nil {
// 			return *memo[length]
// 		}
// 		var count = 0
// 		if length >= low && length <= high {
// 			count++
// 		}
// 		if length+zero <= high {
// 			count = (count + dp(length+zero)) % MOD
// 		}
// 		if length+one <= high {
// 			count = (count + dp(length+one)) % MOD
// 		}
// 		memo[length] = &count
// 		return count
// 	}
// 	return dp(0)
// }
