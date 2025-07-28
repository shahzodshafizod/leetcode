package stacks

// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

func maxDepth(s string) int {
	depth, counter := 0, 0

	for _, c := range s {
		switch c {
		case '(':
			counter++
			depth = max(depth, counter)
		case ')':
			counter--
		}
	}

	return depth
}

// func maxDepth(s string) int {
// 	var depth = 0
// 	var stack = design.NewStack[rune]()
// 	for _, c := range s {
// 		if c == '(' {
// 			stack.Push(c)
// 			depth = max(depth, stack.Size())
// 		} else if c == ')' && stack.Peek() == '(' {
// 			stack.Pop()
// 		}
// 	}
// 	if !stack.Empty() {
// 		depth = 0
// 	}
// 	return depth
// }
