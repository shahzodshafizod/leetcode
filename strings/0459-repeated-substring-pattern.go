package strings

// https://leetcode.com/problems/repeated-substring-pattern/

// Approach: KMP (Knuth-Morris-Pratt) Algorithm
// Time: O(n)
// Space: O(n)
func repeatedSubstringPattern(s string) bool {
	// return strings.Contains((s + s)[1:2*len(s)-1], s)
	var n = len(s)
	var lps = make([]int, n)
	var preflen = 0
	for idx := 1; idx < n; idx++ {
		if s[idx] == s[preflen] {
			preflen++
			lps[idx] = preflen
		} else if preflen > 0 {
			preflen = lps[preflen-1]
			idx--
		}
	}
	return lps[n-1] != 0 && lps[n-1]%(n-lps[n-1]) == 0
}
