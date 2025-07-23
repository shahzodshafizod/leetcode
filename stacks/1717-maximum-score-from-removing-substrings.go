package stacks

// https://leetcode.com/problems/maximum-score-from-removing-substrings/

// Approach 2: Greedy (Counting)
// Time: O(n)
// Space: O(1)
func maximumGain(s string, x int, y int) int {
	a, b := 'a', 'b'
	if x < y {
		x, y = y, x
		a, b = b, a
	}
	score, acnt, bcnt := 0, 0, 0
	for _, c := range s {
		switch c {
		case a:
			acnt++
		case b:
			if acnt > 0 {
				score += x
				acnt--
			} else {
				bcnt++
			}
		default:
			score += min(acnt, bcnt) * y
			acnt, bcnt = 0, 0
		}
	}
	score += min(acnt, bcnt) * y
	return score
}

// // Approach #1: Stack
// // Time: O(2n) = O(n)
// // Space: O(n)
// func maximumGain(s string, x int, y int) int {
// 	bayts := []byte(s)
// 	size := len(bayts)
// 	getScore := func(a byte, b byte, score int) int {
// 		res, ptr := 0, 0
// 		for i := 0; i < size; i++ {
// 			if ptr > 0 && bayts[ptr-1] == a && bayts[i] == b {
// 				res += score
// 				ptr--
// 			} else {
// 				bayts[ptr] = bayts[i]
// 				ptr++
// 			}
// 		}
// 		size = ptr
// 		return res
// 	}
// 	var a, b byte = 'a', 'b'
// 	if x < y {
// 		x, y = y, x
// 		a, b = b, a
// 	}
// 	return getScore(a, b, x) + getScore(b, a, y)
// }
