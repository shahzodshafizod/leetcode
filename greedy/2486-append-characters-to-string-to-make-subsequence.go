package greedy

// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/

// time: O(len(s))
// space: O(1)
func appendCharacters(s string, t string) int {
	var tlen, slen = len(t), len(s)
	var ti = 0
	for si := 0; ti < tlen && si < slen; si++ {
		if t[ti] == s[si] {
			ti++
		}
	}
	return tlen - ti
}
