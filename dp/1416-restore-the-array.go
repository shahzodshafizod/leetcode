package dp

// https://leetcode.com/problems/restore-the-array/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(n log k)
// Space: O(n)
func numberOfArrays(s string, k int) int {
	const mod = int(1e9) + 7

	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1

	var num, j int

	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}

		num = 0
		for j = i; j < n; j++ {
			num = num*10 + int(s[j]-'0')
			if num > k {
				break
			}

			dp[i] = (dp[i] + dp[j+1]) % mod
		}
	}

	return dp[0]
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(n log k)
// // Space: O(n)
// func numberOfArrays(s string, k int) int {
// 	const mod = int(1e9) + 7

// 	n := len(s)
// 	memo := make([]*int, n)

// 	var dfs func(i int) int

// 	dfs = func(i int) int {
// 		if i == n {
// 			return 1
// 		}

// 		if s[i] == '0' {
// 			return 0
// 		}

// 		if memo[i] != nil {
// 			return *memo[i]
// 		}

// 		var ans, num int
// 		for j := i; j < n; j++ {
// 			num = num*10 + int(s[j]-'0')
// 			if num > k {
// 				break
// 			}

// 			ans = (ans + dfs(j+1)) % mod
// 		}

// 		memo[i] = &ans

// 		return ans
// 	}

// 	return dfs(0)
// }
