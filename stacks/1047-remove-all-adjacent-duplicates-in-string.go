package stacks

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

// Approach: Stack + Two Pointers
// Time: O(n)
// Space: O(n)
func removeDuplicates(s string) string {
	var stack = make([]rune, len(s))
	var size = 0
	for _, c := range s {
		if size > 0 && stack[size-1] == c {
			size--
		} else {
			stack[size] = c
			size++
		}
	}
	return string(stack[:size])
}
