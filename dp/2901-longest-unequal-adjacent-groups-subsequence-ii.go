package dp

// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	check := func(i int, j int) bool {
		if len(words[i]) != len(words[j]) {
			return false
		}

		count := 0
		for k := len(words[i]) - 1; k >= 0 && count <= 1; k-- {
			if words[i][k] != words[j][k] {
				count++
			}
		}

		return count == 1
	}
	n := len(words)
	dp := make([]int, n)
	next := make([]*int, n)
	best := n - 1

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if groups[i] != groups[j] && dp[j]+1 > dp[i] && check(i, j) {
				dp[i] = dp[j] + 1
				next[i] = new(int)
				*next[i] = j
			}
		}

		if dp[i] > dp[best] {
			best = i
		}
	}

	var subsequence []string
	for curr := &best; curr != nil; curr = next[*curr] {
		subsequence = append(subsequence, words[*curr])
	}

	return subsequence
}
