package prefixsums

import "math/bits"

// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
func countPalindromicSubsequence(s string) int {
	n := len(s)
	right := make([]int, n)
	mask := 0
	for idx := n - 1; idx >= 0; idx-- {
		mask |= 1 << int(s[idx]-'a')
		right[idx] = mask
	}
	left := 0
	var masks [26]int
	var key int
	for idx := 0; idx < n-1; idx++ {
		key = int(s[idx] - 'a')
		masks[key] |= left & right[idx+1]
		left |= 1 << key
	}
	count := 0
	for _, mask := range masks {
		count += bits.OnesCount(uint(mask))
	}
	return count
}
