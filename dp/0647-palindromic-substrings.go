package dp

// https://leetcode.com/problems/palindromic-substrings/

func countSubstrings(s string) int {
	n := len(s)
	expand := func(left int, right int) int {
		count := 0
		for left >= 0 && right < n && s[left] == s[right] {
			count++
			left--
			right++
		}

		return count
	}
	count := 0

	for i := range n {
		count += expand(i, i)
		if i+1 < n {
			count += expand(i, i+1)
		}
	}

	return count
}
