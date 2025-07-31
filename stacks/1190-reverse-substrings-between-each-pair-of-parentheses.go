package stacks

// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/

// Approach 2: Wormhole Teleportation technique
func reverseParentheses(s string) string {
	pairs := make(map[int]int)
	stack := make([]int, 0)

	for idx := range s {
		switch s[idx] {
		case '(':
			stack = append(stack, idx)
		case ')':
			opening := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pairs[opening] = idx
			pairs[idx] = opening
		}
	}

	idx, dir := 0, 1
	n := len(s)
	result := make([]byte, 0)

	for idx < n {
		if s[idx] == '(' || s[idx] == ')' {
			idx = pairs[idx]
			dir *= -1
		} else {
			result = append(result, s[idx])
		}

		idx += dir
	}

	return string(result)
}

// // Approach 1: Straightforward Way (Brute Force)
// func reverseParentheses(s string) string {
// 	var stack = make([]rune, 0)
// 	for _, r := range s {
// 		if r == ')' {
// 			var portion = make([]rune, 0)
// 			var len = len(stack)
// 			for stack[len-1] != '(' {
// 				portion = append(portion, stack[len-1])
// 				len--
// 			}
// 			stack = append(stack[:len-1], portion...)
// 		} else {
// 			stack = append(stack, r)
// 		}
// 	}
// 	return string(stack)
// }
