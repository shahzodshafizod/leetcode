package stacks

/*
Problem:
Given a string containing only parentheses, determine if it is valid.
The string is valid if all parentheses close.

Step 1: Verify the constraints
	- Does an empty string count as valid?
		: Yes, return true if empty string.
Step 2: Write out some test cases
	- "" -> true
	- "{([])}" -> true
	- "{([]" -> false
	- "{[(])}" -> false
	- "{[]()}" -> true
*/

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]byte, 0)
	length := 0
	parentheses := map[byte]byte{'}': '{', ')': '(', ']': '['}

	for i := len(s) - 1; i >= 0; i-- {
		if length > 0 && s[i] == parentheses[stack[length-1]] {
			length--
		} else {
			if len(stack) > length {
				stack[length] = s[i]
			} else {
				stack = append(stack, s[i])
			}

			length++
			if stack[length-1] == parentheses[s[i]] {
				return false
			}
		}
	}

	return length == 0
}

// func isValid(s string) bool {
// 	var stack = make([]byte, 0)
// 	length := 0
// 	for i := len(s) - 1; i >= 0; i-- {
// 		remove := false
// 		if length > 0 {
// 			correctParenthes := false
// 			switch s[i] {
// 			case '{':
// 				remove = stack[length-1] == '}'
// 			case '(':
// 				remove = stack[length-1] == ')'
// 			case '[':
// 				remove = stack[length-1] == ']'
// 			default:
// 				correctParenthes = true
// 			}
// 			if !remove && !correctParenthes {
// 				return false
// 			}
// 		}
// 		if remove {
// 			length--
// 		} else {
// 			if len(stack) > length {
// 				stack[length] = s[i]
// 			} else {
// 				stack = append(stack, s[i])
// 			}
// 			length++
// 		}
// 	}
// 	return length == 0
// }
