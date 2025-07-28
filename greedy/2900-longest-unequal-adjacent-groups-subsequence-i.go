package greedy

// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/

func getLongestSubsequence(words []string, groups []int) []string {
	var ans []string

	for i, n := 0, len(words); i < n; i++ {
		if i == 0 || groups[i] != groups[i-1] {
			ans = append(ans, words[i])
		}
	}

	return ans
}
