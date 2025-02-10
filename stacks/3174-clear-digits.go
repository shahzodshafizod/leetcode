package stacks

// https://leetcode.com/problems/clear-digits/

func clearDigits(s string) string {
	var n = len(s)
	var stack = make([]byte, n)
	var size = 0
	for idx := 0; idx < n; idx++ {
		if '0' <= s[idx] && s[idx] <= '9' {
			size--
		} else {
			stack[size] = s[idx]
			size++
		}
	}
	return string(stack[:size])
}
