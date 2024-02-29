package dp

// https://leetcode.com/problems/palindromic-substrings/

func countSubstrings(s string) int {
	var n = len(s)
	var expand = func(left int, right int) int {
		var count = 0
		for left >= 0 && right < n && s[left] == s[right] {
			count++
			left--
			right++
		}
		return count
	}
	var count = 0
	for i := 0; i < n; i++ {
		count += expand(i, i)
		if i+1 < n {
			count += expand(i, i+1)
		}
	}
	return count
}
