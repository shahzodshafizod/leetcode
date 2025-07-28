package dp

// https://leetcode.com/problems/longest-ideal-subsequence/

// Tabulation
func longestIdealString(s string, k int) int {
	var dp [26]int

	var current, letter int

	length := 0

	for idx := range s {
		current = int(s[idx] - 'a')
		dp[current]++

		for letter = max(0, current-k); letter < current; letter++ {
			dp[current] = max(dp[current], 1+dp[letter])
		}

		for letter = min(25, current+k); letter > current; letter-- {
			dp[current] = max(dp[current], 1+dp[letter])
		}

		length = max(length, dp[current])
	}

	return length
}

// // Memoization
// func longestIdealString(s string, k int) int {
// 	var n = len(s)
// 	var cache = make([][26]int, n)
// 	var dp func(idx int, prev int) int
// 	dp = func(idx int, prev int) int {
// 		if idx == n {
// 			return 0
// 		}
// 		if prev >= 0 && cache[idx][prev] != 0 {
// 			return cache[idx][prev]
// 		}
// 		var curr = int(s[idx] - 'a')
// 		var length = dp(idx+1, prev)
// 		if prev == -1 || int(math.Abs(float64(curr-prev))) <= k {
// 			length = max(length, 1+dp(idx+1, curr))
// 		}
// 		if prev >= 0 {
// 			cache[idx][prev] = length
// 		}
// 		return length
// 	}
// 	return dp(0, -1)
// }
