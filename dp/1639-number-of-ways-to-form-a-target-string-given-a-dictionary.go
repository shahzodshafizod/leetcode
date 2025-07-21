package dp

// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/

// Approach: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(W*T + W*w), W=len(words[i]), T=len(target), w=len(words)
// Space: O(T)
func numWays1639(words []string, target string) int {
	tlen, wlen := len(target), len(words[0])
	count := make([][26]int, wlen)
	for _, word := range words {
		for wid := range word {
			count[wid][int(word[wid]-'a')]++
		}
	}
	curr, next := make([]int, tlen+1), make([]int, tlen+1)
	curr[tlen], next[tlen] = 1, 1
	const MOD int = 1e9 + 7
	for wid := wlen - 1; wid >= 0; wid-- {
		curr, next = next, curr
		for tid := tlen - 1; tid >= 0; tid-- {
			curr[tid] = (next[tid] +
				count[wid][int(target[tid]-'a')]*next[tid+1]) % MOD
		}
	}
	return curr[0]
}

// // Approach: Bottom-Up Dynamic Programming
// // Time: O(W*T + W*w), W=len(words[i]), T=len(target), w=len(words)
// // Space: O(W*T)
// func numWays1639(words []string, target string) int {
// 	var tlen, wlen = len(target), len(words[0])
// 	var count = make([][26]int, wlen)
// 	for _, word := range words {
// 		for wid := range word {
// 			count[wid][int(word[wid]-'a')]++
// 		}
// 	}
// 	var dp = make([][]int, wlen+1)
// 	for idx := range dp {
// 		dp[idx] = make([]int, tlen+1)
// 		dp[idx][tlen] = 1
// 	}
// 	const MOD int = 1e9 + 7
// 	for wid := wlen - 1; wid >= 0; wid-- {
// 		for tid := tlen - 1; tid >= 0; tid-- {
// 			dp[wid][tid] = (dp[wid+1][tid] +
// 				count[wid][int(target[tid]-'a')]*dp[wid+1][tid+1]) % MOD
// 		}
// 	}
// 	return dp[0][0]
// }

// // Approach: Top-Down Dynamic Programming
// // Time: O(W*T + W*w), W=len(words[i]), T=len(target), w=len(words)
// // Space: O(W*T)
// func numWays1639(words []string, target string) int {
// 	const MOD int = 1e9 + 7
// 	var wlen, tlen = len(words[0]), len(target)
// 	if tlen > wlen {
// 		return 0
// 	}
// 	var count = make([][26]int, wlen)
// 	for _, word := range words {
// 		for idx := range word {
// 			count[idx][int(word[idx]-'a')]++
// 		}
// 	}
// 	var memo = make([][]*int, wlen)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, tlen)
// 	}
// 	var dp func(wid int, tid int) int
// 	dp = func(wid int, tid int) int {
// 		if tid == tlen {
// 			return 1
// 		}
// 		if wid == wlen {
// 			return 0
// 		}
// 		if memo[wid][tid] != nil {
// 			return *memo[wid][tid]
// 		}
// 		// decision to skip
// 		var ways = dp(wid+1, tid)
// 		// decision to include
// 		if count[wid][int(target[tid]-'a')] > 0 {
// 			ways = (ways +
// 				count[wid][int(target[tid]-'a')]*dp(wid+1, tid+1)) % MOD
// 		}
// 		memo[wid][tid] = &ways
// 		return ways
// 	}
// 	return dp(0, 0)
// }
