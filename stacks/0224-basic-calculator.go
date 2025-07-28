package stacks

/*
1. digit: it should be one digit from the current number
2. '+': number is over, we can add the previous number and start a new number
3. '-': same as above
4. '(': push the previous result and the sign into the stack, set result to 0,
	just calculate the new result within the parenthesis.
5. ')': pop out the top two numbers from stack, first one is the sign before this pair of parenthesis,
	second is the temporary result before this pair of parenthesis. We add them together.
6. Finally if there is only one number, from the above solution,
	we haven't add the number to the result, so we do a check see if the number is zero.
*/

// https://leetcode.com/problems/basic-calculator/

func calculate(s string) int {
	var (
		result = 0
		sign   = 1
		number = 0
	)

	stack := make([][2]int, 0)

	for _, r := range s {
		switch r {
		case '+', '-':
			result += sign * number
			number = 0

			if r == '+' {
				sign = 1
			} else {
				sign = -1
			}
		case '(':
			stack = append(stack, [2]int{sign, result})
			sign = 1
			result = 0
		case ')':
			result += sign * number
			sign = 1
			number = 0
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = result*prev[0] + prev[1]
		case ' ':
			if number != 0 {
				result += sign * number
				number = 0
			}
		default:
			number = number*10 + int(r-'0')
		}
	}

	return result + sign*number
}

// func calculate(s string) int {
// 	var index = 0
// 	return calculateRecur(s, &index, len(s))
// }

// func calculateRecur(s string, index *int, n int) int {
// 	if *index >= n {
// 		return 0
// 	}
// 	var operations = NewStack[byte]()
// 	var numbers = NewStack[int]()
// 	for *index < n {
// 		r := s[*index]
// 		switch r {
// 		case '+', '-':
// 			operations.Push(r)
// 			if numbers.Empty() {
// 				numbers.Push(0) // very useful for unary minuses
// 			}
// 		case '(':
// 			(*index)++
// 			number := calculateRecur(s, index, n)
// 			if !operations.Empty() {
// 				first := numbers.Pop()
// 				switch operations.Pop() {
// 				case '+':
// 					number += first
// 				case '-':
// 					number = first - number
// 				}
// 			}
// 			numbers.Push(number)
// 		case ')':
// 			return numbers.Pop()
// 		case ' ':
// 			if numbers.Size() > operations.Size() {
// 				operations.Push('+')
// 			}
// 		default:
// 			number := int(r - '0')
// 			for *index+1 < n && s[*index+1] >= '0' && s[*index+1] <= '9' {
// 				(*index)++
// 				number = number*10 + int(s[*index]-'0')
// 			}
// 			if !operations.Empty() {
// 				first := numbers.Pop()
// 				switch operations.Pop() {
// 				case '+':
// 					number += first
// 				case '-':
// 					number = first - number
// 				}
// 			}
// 			numbers.Push(number)
// 		}
// 		(*index)++
// 	}
// 	return numbers.Pop()
// }
