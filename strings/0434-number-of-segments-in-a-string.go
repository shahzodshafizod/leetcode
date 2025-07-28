package strings

// https://leetcode.com/problems/number-of-segments-in-a-string/

func countSegments(s string) int {
	count := 0

	for i, n := 0, len(s); i < n; i++ {
		if s[i] != ' ' && (i == 0 || s[i-1] == ' ') {
			count++
		}
	}

	return count
}
