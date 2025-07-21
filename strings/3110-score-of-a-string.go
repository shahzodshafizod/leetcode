package strings

// https://leetcode.com/problems/score-of-a-string/

func scoreOfString(s string) int {
	score := 0
	for i, n := 1, len(s); i < n; i++ {
		if s[i] > s[i-1] {
			score += int(s[i] - s[i-1])
		} else {
			score += int(s[i-1] - s[i])
		}
	}
	return score
}
