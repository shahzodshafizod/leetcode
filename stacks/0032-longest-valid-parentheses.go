package stacks

// https://leetcode.com/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	length := func(open byte, from int, to int, factor int) int {
		length, opens, closes := 0, 0, 0

		for idx := from; idx != to; idx += factor {
			if s[idx] == open {
				opens++
			} else {
				closes++
			}

			if closes > opens {
				opens, closes = 0, 0
			} else if opens == closes {
				length = max(length, opens+closes)
			}
		}

		return length
	}

	return max(length('(', 0, len(s), 1), length(')', len(s)-1, -1, -1))
}

// func longestValidParentheses(s string) int {
// 	var sb = []byte(s)
// 	var stack = make([]int, 0)
// 	for idx := range sb {
// 		switch s[idx] {
// 		case '(':
// 			stack = append(stack, idx)
// 		case ')':
// 			if len(stack) > 0 {
// 				sb[stack[len(stack)-1]] = '.'
// 				sb[idx] = '.'
// 				stack = stack[:len(stack)-1]
// 			}
// 		}
// 	}
// 	var length, count = 0, 0
// 	for _, c := range sb {
// 		switch c {
// 		case '.':
// 			count++
// 		default:
// 			count = 0
// 		}
// 		length = max(length, count)
// 	}
// 	return length
// }

// func longestValidParentheses(s string) int {
// 	var sb = []byte(s)
// 	var n = len(sb)
// 	var dp func(start int, end int) int
// 	dp = func(start int, end int) int {
// 		if end < 0 || end >= n {
// 			return -1
// 		}
// 		switch sb[end] {
// 		case '(':
// 			return dp(start, dp(end, end+1))
// 		case ')':
// 			if start >= 0 {
// 				sb[start] = '.'
// 				sb[end] = '.'
// 				return end + 1
// 			}
// 			dp(start, end+1)
// 			return end
// 		default:
// 			return dp(start, end+1)
// 		}
// 	}
// 	dp(-1, 0)
// 	var length, count = 0, 0
// 	for _, b := range sb {
// 		switch b {
// 		case '.':
// 			count++
// 		default:
// 			count = 0
// 		}
// 		length = max(length, count)
// 	}
// 	return length
// }
