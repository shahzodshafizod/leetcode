package stacks

// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/

// Approach: Stack
// Time: O(N)
// Space: O(N)
func minLength(s string) int {
	stack := make([]rune, len(s))
	idx := -1

	for _, c := range s {
		if idx >= 0 && (c == 'B' || c == 'D') && rune(stack[idx]+1) == c {
			idx--
		} else {
			idx++
			stack[idx] = c
		}
	}

	return idx + 1
}
