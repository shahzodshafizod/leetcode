package stacks

// https://leetcode.com/problems/baseball-game/

// Approach: Optimized - Stack simulation
// Use stack for natural push/pop operations
// Process each operation: +, D, C, or number
// N = number of operations
// Time: O(N) - process each operation once, final sum is O(N)
// Space: O(N) - stack stores all valid scores
func calPoints(operations []string) int {
	stack := make([]int, 0)

	for _, op := range operations {
		n := len(stack)

		switch op {
		case "+":
			// Sum of previous two scores
			stack = append(stack, stack[n-1]+stack[n-2])
		case "D":
			// Double of previous score
			stack = append(stack, stack[n-1]*2)
		case "C":
			// Remove previous score
			stack = stack[:n-1]
		default:
			// It's a number (parse and add)
			var num int
			// Manual parsing to handle negative numbers
			sign := 1
			idx := 0

			if op[0] == '-' {
				sign = -1
				idx = 1
			}

			for ; idx < len(op); idx++ {
				num = num*10 + int(op[idx]-'0')
			}

			stack = append(stack, sign*num)
		}
	}

	// Calculate sum of all scores
	sum := 0
	for _, score := range stack {
		sum += score
	}

	return sum
}
