package dp

// https://leetcode.com/problems/find-the-original-typed-string-ii/

// Approach: Bottom-Up Dynamic Programming (Tabulation) + Prefix Sum
// Time: O(n + m*k), n=len(word), m=len(groups)
// Space: O(k)
func possibleStringCount(word string, k int) int {
	const MOD int = 1e9 + 7
	var groups []int
	var count, total = 1, 1
	var n = len(word)
	for idx := 1; idx < n; idx++ {
		if word[idx-1] != word[idx] {
			groups = append(groups, count)
			total = (total * count) % MOD
			count = 0
		}
		count++
	}
	groups = append(groups, count)
	total = (total * count) % MOD
	var m = len(groups)
	if k <= m {
		return total
	}
	var dp = make([]int, k)
	dp[0] = 1
	var presum int
	for _, count := range groups {
		var newDP = make([]int, k)
		presum = 0
		for idx := 1; idx < k; idx++ {
			presum = (presum + dp[idx-1]) % MOD
			if idx > count {
				presum = (presum - dp[idx-count-1] + MOD) % MOD
			}
			newDP[idx] = presum
		}
		dp = newDP
	}
	for _, count := range dp {
		total = (total - count + MOD) % MOD
	}
	return total
}
