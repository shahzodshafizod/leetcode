package dp

// https://leetcode.com/problems/distinct-subsequences/

// Approach #3: Bottom-Up Dynamic Programming (Space Optimized)
// Time: O(ns * nt), ns=len(s), nt=len(t)
// Space: O(nt)
func numDistinct(s string, t string) int {
	var ns, nt = len(s), len(t)
	var dp = make([]int, nt+1)
	dp[0] = 1
	for is := 1; is <= ns; is++ {
		for it := nt; it > 0; it-- {
			if s[is-1] == t[it-1] {
				dp[it] += dp[it-1]
			}
		}
	}
	return dp[nt]
}

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(ns * nt), ns=len(s), nt=len(t)
// // Space: O(ns * nt)
// func numDistinct(s string, t string) int {
// 	var ns, nt = len(s), len(t)
// 	var dp = make([][]int, ns+1)
// 	for idx := range dp {
// 		dp[idx] = make([]int, nt+1)
// 		dp[idx][0] = 1
// 	}
// 	for is := 1; is <= ns; is++ {
// 		for it := 1; it <= nt; it++ {
// 			if s[is-1] == t[it-1] {
// 				dp[is][it] = dp[is-1][it] + dp[is-1][it-1]
// 			} else {
// 				dp[is][it] = dp[is-1][it]
// 			}
// 		}
// 	}
// 	return dp[ns][nt]
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(ns * nt), ns=len(s), nt=len(t)
// // Space: O(ns * nt)
// func numDistinct(s string, t string) int {
// 	var ns, nt = len(s), len(t)
// 	var dp = make([][]*int, ns)
// 	for idx := range dp {
// 		dp[idx] = make([]*int, nt)
// 	}
// 	var dfs func(is int, it int) int
// 	dfs = func(is int, it int) int {
// 		if it == nt {
// 			return 1
// 		}
// 		if is == ns {
// 			return 0
// 		}
// 		if dp[is][it] != nil {
// 			return *dp[is][it]
// 		}
// 		dp[is][it] = new(int)
// 		if s[is] != t[it] {
// 			*dp[is][it] = dfs(is+1, it)
// 		} else {
// 			*dp[is][it] = dfs(is+1, it) + dfs(is+1, it+1)
// 		}
// 		return *dp[is][it]
// 	}
// 	return dfs(0, 0)
// }
