package greedy

import "strings"

// https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/

// Approach: Greedy with Stack
// Time: O(n)
// Space: O(n)
func clearStars(s string) string {
	var stack [26][]int
	stars := 0
	starred := []byte(s)
	for idx, c := range s {
		if c == '*' {
			stars++
			for k := 0; k < 26; k++ {
				if n := len(stack[k]); n != 0 {
					starred[stack[k][n-1]] = '*'
					stack[k] = stack[k][:n-1]
					break
				}
			}
		} else {
			stack[int(c-'a')] = append(stack[int(c-'a')], idx)
		}
	}
	if stars == 0 {
		return s
	}
	var cleared strings.Builder
	cleared.Grow(len(s) - stars*2)
	for _, c := range starred {
		if c != '*' {
			cleared.WriteByte(c)
		}
	}
	return cleared.String()
}
