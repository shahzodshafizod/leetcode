package prefixsums

import "strings"

// https://leetcode.com/problems/maximum-score-after-splitting-a-string/

func maxScore(s string) int {
	zeroes, ones := 0, strings.Count(s, "1")
	score := 0

	n := len(s)
	for idx := range n - 1 {
		if s[idx] == '0' {
			zeroes++
		} else {
			ones--
		}

		score = max(score, zeroes+ones)
	}

	return score
}
