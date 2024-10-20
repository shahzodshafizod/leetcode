package stacks

// https://leetcode.com/problems/parsing-a-boolean-expression/

// Approach: Iterative (Using Stack)
// Time: O(N)
// Space: O(N)
func parseBoolExpr(expression string) bool {
	var stack = make([]rune, len(expression))
	var sid = -1
	for _, c := range expression {
		if c == ')' {
			var value rune
			var hasTrue, hasFalse = false, false
			for stack[sid] != '(' {
				value = stack[sid]
				sid--
				hasTrue = hasTrue || value == 't'
				hasFalse = hasFalse || value == 'f'
			}
			sid-- // skip ')'
			// sid++ // (for a new element) we'll replace the operator
			var oper = stack[sid]
			if oper == '!' && value == 'f' ||
				oper == '&' && !hasFalse ||
				oper == '|' && hasTrue {
				stack[sid] = 't'
			} else {
				stack[sid] = 'f'
			}
		} else if c != ',' {
			sid++
			stack[sid] = c
		}
	}
	return stack[sid] == 't'
}

// // Approach: Recursive
// // Time: O(N)
// // Space: O(N)
// func parseBoolExpr(expression string) bool {
// 	var idx = 0
// 	var parse func() bool
// 	parse = func() bool {
// 		var c = expression[idx]
// 		idx++
// 		switch c {
// 		case 't':
// 			return true
// 		case 'f':
// 			return false
// 		case '!':
// 			idx++ // skipping the openning parenthesis
// 			result := !parse()
// 			idx++ // skipping the closing parenthesis
// 			return result
// 		}
// 		idx++ // skip '('
// 		var and, or = true, false
// 		for expression[idx] != ')' {
// 			if expression[idx] == ',' {
// 				idx++
// 			} else if c == '&' {
// 				and = parse() && and // parse() should be first!
// 			} else {
// 				or = parse() || or // parse() should be first!
// 			}
// 		}
// 		idx++ // skip ')'
// 		if c == '&' {
// 			return and
// 		}
// 		return or
// 	}
// 	return parse()
// }
