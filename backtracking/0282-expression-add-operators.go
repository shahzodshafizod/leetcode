package backtracking

import (
	"strconv"
)

// https://leetcode.com/problems/expression-add-operators/

// Approach 2: Optimized Backtracking with evaluation during construction
// Instead of evaluating at the end, we maintain the current value and last operand
// This allows us to handle multiplication correctly by reversing the effect of the last operation
// Time: O(4^n) where n is the length of num
//   - At each position, we try 4 possibilities (extend number or add +, -, *)
//   - No need for separate evaluation step
//
// Space: O(n) for recursion depth and expression string
func addOperators(num string, target int) []string {
	result := []string{}
	n := len(num)

	var backtrack func(index int, expr string, currVal int64, lastOperand int64)

	backtrack = func(index int, expr string, currVal int64, lastOperand int64) {
		// currVal: current evaluation result
		// lastOperand: the value of the last number we added/subtracted
		//              (needed to handle multiplication correctly)
		if index == n {
			if currVal == int64(target) {
				result = append(result, expr)
			}

			return
		}

		// Try all possible lengths of the next number (1 to n-index digits)
		for length := 1; length <= n-index; length++ {
			numStr := num[index : index+length]

			// Skip numbers with leading zeros (except "0" itself)
			if len(numStr) > 1 && numStr[0] == '0' {
				break
			}

			currentNum, _ := strconv.ParseInt(numStr, 10, 64)

			if index == 0 {
				// First number, no operator needed
				backtrack(index+length, numStr, currentNum, currentNum)
			} else {
				// Addition: simply add to current value
				backtrack(
					index+length,
					expr+"+"+numStr,
					currVal+currentNum,
					currentNum,
				)

				// Subtraction: subtract from current value
				backtrack(
					index+length,
					expr+"-"+numStr,
					currVal-currentNum,
					-currentNum,
				)

				// Multiplication: reverse the last operation and apply multiplication
				// Example: if we had "1+2" (currVal=3, lastOperand=2) and multiply by 3:
				// We need: 1+2*3 = 1+6 = 7
				// So: currVal - lastOperand + (lastOperand * currentNum)
				//     = 3 - 2 + (2 * 3) = 1 + 6 = 7
				backtrack(
					index+length,
					expr+"*"+numStr,
					currVal-lastOperand+lastOperand*currentNum,
					lastOperand*currentNum,
				)
			}
		}
	}

	backtrack(0, "", 0, 0)

	return result
}

// // Approach 1: Brute Force - Generate all possible expressions and evaluate each
// // We try every possible way to insert operators and multi-digit numbers
// // Time: O(4^n * n) where n is the length of num
// //   - At each position, we can choose to add an operator or continue building a number
// //   - Roughly 4 choices at each step: +num, -num, *num, or extend current number
// //   - Each expression evaluation takes O(n) time
// //
// // Space: O(n) for recursion depth and expression string
// func addOperators(num string, target int) []string {
// 	result := []string{}
// 	n := len(num)

// 	// Helper function to evaluate an expression string
// 	evaluate := func(expr string) int {
// 		// Tokenize the expression
// 		tokens := []string{}
// 		currentNum := ""

// 		for i := range len(expr) {
// 			if (expr[i] >= '0' && expr[i] <= '9') || (expr[i] == '-' && i == 0) {
// 				currentNum += string(expr[i])
// 			} else {
// 				tokens = append(tokens, currentNum)
// 				tokens = append(tokens, string(expr[i]))
// 				currentNum = ""
// 			}
// 		}

// 		if currentNum != "" {
// 			tokens = append(tokens, currentNum)
// 		}

// 		// First pass: handle multiplication
// 		i := 0
// 		for i < len(tokens) {
// 			if i > 0 && tokens[i] == "*" {
// 				left, _ := strconv.Atoi(tokens[i-1])
// 				right, _ := strconv.Atoi(tokens[i+1])
// 				product := strconv.Itoa(left * right)
// 				tokens = append(append(tokens[:i-1], product), tokens[i+2:]...)
// 			} else {
// 				i++
// 			}
// 		}

// 		// Second pass: handle addition and subtraction
// 		result, _ := strconv.Atoi(tokens[0])

// 		i = 1
// 		for i < len(tokens) {
// 			op := tokens[i]

// 			num, _ := strconv.Atoi(tokens[i+1])
// 			switch op {
// 			case "+":
// 				result += num
// 			case "-":
// 				result -= num
// 			}

// 			i += 2
// 		}

// 		return result
// 	}

// 	var backtrack func(index int, expr string)

// 	backtrack = func(index int, expr string) {
// 		if index == n {
// 			if evaluate(expr) == target {
// 				result = append(result, expr)
// 			}

// 			return
// 		}

// 		// Try all possible lengths of the next number (1 to n-index digits)
// 		for length := 1; length <= n-index; length++ {
// 			numStr := num[index : index+length]

// 			// Skip numbers with leading zeros (except "0" itself)
// 			if len(numStr) > 1 && numStr[0] == '0' {
// 				break
// 			}

// 			if index == 0 {
// 				// First number, no operator needed
// 				backtrack(index+length, numStr)
// 			} else {
// 				// Try each operator
// 				backtrack(index+length, expr+"+"+numStr)
// 				backtrack(index+length, expr+"-"+numStr)
// 				backtrack(index+length, expr+"*"+numStr)
// 			}
// 		}
// 	}

// 	backtrack(0, "")

// 	return result
// }
