package greedy

import "math/bits"

// https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/

func longestSubsequence(s string, k int) int {
	count, bits := 0, bits.Len(uint(k))

	for i, n := 0, len(s); i < n; i++ {
		if s[n-1-i] == '0' {
			count++
		} else if i < bits && k >= (1<<i) {
			k -= 1 << i
			count++
		}
	}

	return count
}
