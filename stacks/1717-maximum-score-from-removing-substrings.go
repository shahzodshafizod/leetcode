package stacks

// https://leetcode.com/problems/maximum-score-from-removing-substrings/

// Approach 2: Greedy Way (Counting)
// time: O(N)
// space: O(1)
func maximumGain(s string, x int, y int) int {
	var a, b = 'a', 'b'
	if x < y {
		x, y = y, x
		a, b = b, a
	}
	var points = 0
	var acount, bcount = 0, 0
	for _, r := range s {
		switch r {
		case a:
			acount++
		case b:
			if acount > 0 {
				acount--
				points += x
			} else {
				bcount++
			}
		default:
			points += min(acount, bcount) * y
			acount, bcount = 0, 0
		}
	}
	points += min(acount, bcount) * y
	return points
}

// // Approach #1: Stack
// // time: O(2*N) = O(N)
// // space: O(2*N) = O(N)
// func maximumGain(s string, x int, y int) int {
// 	var a, b = 'a', 'b'
// 	if y > x {
// 		x, y = y, x
// 		a, b = b, a
// 	}
// 	var points = 0
// 	var stack = design.NewStack[rune]()
// 	for _, r := range s {
// 		if !stack.Empty() && stack.Top() == a && r == b {
// 			stack.Pop()
// 			points += x
// 		} else {
// 			stack.Push(r)
// 		}
// 	}
// 	var stack2 = design.NewStack[rune]()
// 	var r rune
// 	for !stack.Empty() {
// 		r = stack.Pop()
// 		if !stack2.Empty() && stack2.Top() == a && r == b {
// 			stack2.Pop()
// 			points += y
// 		} else {
// 			stack2.Push(r)
// 		}
// 	}
// 	return points
// }
