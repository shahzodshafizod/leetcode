package backtracking

// https://leetcode.com/problems/remove-invalid-parentheses/

// Approach 2: Backtracking (Optimized)
// First calculate how many '(' and ')' need to be removed
// Then use backtracking to generate all valid combinations
// Time: O(2^n) where n is the length of string
// Space: O(n) for recursion depth
func removeInvalidParentheses(s string) []string {
	// Step 1: Count how many left and right parentheses to remove
	leftToRemove := 0
	rightToRemove := 0

	for _, char := range s {
		switch char {
		case '(':
			leftToRemove++
		case ')':
			if leftToRemove > 0 {
				leftToRemove--
			} else {
				rightToRemove++
			}
		default:
		}
	}

	resultSet := make(map[string]bool)

	var backtrack func(index int, leftCount int, rightCount int, leftRem int, rightRem int, expr string)

	backtrack = func(index int, leftCount int, rightCount int, leftRem int, rightRem int, expr string) {
		// index: current position in string
		// leftCount: number of '(' in current expression
		// rightCount: number of ')' in current expression
		// leftRem: number of '(' still need to remove
		// rightRem: number of ')' still need to remove
		// expr: current expression being built

		// Base case: reached end of string
		if index == len(s) {
			// Valid if we removed all invalid parentheses and counts match
			if leftRem == 0 && rightRem == 0 && leftCount == rightCount {
				resultSet[expr] = true
			}

			return
		}

		char := s[index]

		// Option 1: Remove current character (if it's a parenthesis and we need to remove)
		if char == '(' && leftRem > 0 {
			backtrack(index+1, leftCount, rightCount, leftRem-1, rightRem, expr)
		} else if char == ')' && rightRem > 0 {
			backtrack(index+1, leftCount, rightCount, leftRem, rightRem-1, expr)
		}

		// Option 2: Keep current character
		switch char {
		case '(':
			backtrack(index+1, leftCount+1, rightCount, leftRem, rightRem, expr+string(char))
		case ')':
			// Only add ')' if it doesn't make expression invalid
			if rightCount < leftCount {
				backtrack(index+1, leftCount, rightCount+1, leftRem, rightRem, expr+string(char))
			}
		default:
			// Regular character, always keep
			backtrack(index+1, leftCount, rightCount, leftRem, rightRem, expr+string(char))
		}
	}

	backtrack(0, 0, 0, leftToRemove, rightToRemove, "")

	result := make([]string, 0, len(resultSet))
	for str := range resultSet {
		result = append(result, str)
	}

	if len(result) == 0 {
		return []string{""}
	}

	return result
}

// // Approach 1: BFS (Breadth-First Search)
// // Generate all possible strings by removing one parenthesis at a time
// // Use BFS to ensure we find the minimum number of removals
// // Time: O(2^n) where n is the length of string (each char can be kept or removed)
// // Space: O(2^n) for storing all possible strings in the queue
// func removeInvalidParentheses(s string) []string {
// 	isValid := func(str string) bool {
// 		count := 0

// 		for _, char := range str {
// 			switch char {
// 			case '(':
// 				count++
// 			case ')':
// 				count--
// 				if count < 0 {
// 					return false
// 				}
// 			}
// 		}

// 		return count == 0
// 	}

// 	result := []string{}
// 	visited := make(map[string]bool)
// 	queue := []string{s}
// 	visited[s] = true
// 	found := false

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		if isValid(current) {
// 			result = append(result, current)
// 			found = true
// 		}

// 		// If we found valid strings at this level, don't go deeper
// 		if found {
// 			continue
// 		}

// 		// Try removing each parenthesis
// 		for i := range len(current) {
// 			if current[i] != '(' && current[i] != ')' {
// 				continue
// 			}

// 			nextStr := current[:i] + current[i+1:]
// 			if !visited[nextStr] {
// 				visited[nextStr] = true
// 				queue = append(queue, nextStr)
// 			}
// 		}
// 	}

// 	if len(result) == 0 {
// 		return []string{""}
// 	}

// 	return result
// }
