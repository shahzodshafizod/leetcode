package twopointers

// https://leetcode.com/problems/reverse-string-ii/

// Approach: Two Pointers
// Time: O(n)
// Space: O(n), for b := []byte(s)
func reverseStr(s string, k int) string {
	b, n := []byte(s), len(s)
	var left, right int
	for start := 0; start < n; start += 2 * k {
		left, right = start, min(n-1, start+k-1)
		for left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}
	return string(b)
}
