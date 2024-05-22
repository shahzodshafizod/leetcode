package backtracking

// https://leetcode.com/problems/palindrome-partitioning/

func partition(s string) [][]string {
	var isPalindrome = func(left int, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	var res = make([][]string, 0)
	var part []string
	var dfs func(sid int)
	dfs = func(sid int) {
		if sid == len(s) {
			res = append(res, append([]string{}, part...))
			return
		}
		for j := sid; j < len(s); j++ {
			if isPalindrome(sid, j) {
				part = append(part, s[sid:j+1])
				dfs(j + 1)
				part = part[:len(part)-1]
			}
		}
	}
	dfs(0)
	return res
}
