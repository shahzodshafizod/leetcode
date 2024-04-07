package stacks

// https://leetcode.com/problems/valid-parenthesis-string/

func checkValidString(s string) bool {
	openMin, openMax := 0, 0
	for idx := range s {
		switch s[idx] {
		case '(':
			openMin++
			openMax++
		case ')':
			openMin--
			openMax--
		default:
			openMin-- // consider * as )
			openMax++ // consider * as (
		}
		if openMin < 0 {
			openMin = 0 // ignore *
		}
		if openMax < 0 {
			return false
		}
	}
	return openMin == 0
}

// func checkValidString(s string) bool {
// 	var valid = func(idx int, limit int, openc byte, factor int) bool {
// 		var opens = make([]int, 0)
// 		var stars = make([]int, 0)
// 		for idx != limit {
// 			switch s[idx] {
// 			case '*':
// 				stars = append(stars, idx)
// 			case openc:
// 				opens = append(opens, idx)
// 			default:
// 				switch {
// 				case len(opens) > 0:
// 					opens = opens[:len(opens)-1]
// 				case len(stars) > 0:
// 					stars = stars[:len(stars)-1]
// 				default:
// 					return false
// 				}
// 			}
// 			idx += factor
// 		}
// 		for len(opens) > 0 && len(stars) > 0 &&
// 			(factor == 1 && stars[len(stars)-1] > opens[len(opens)-1] ||
// 				factor == -1 && stars[len(stars)-1] < opens[len(opens)-1]) {
// 			opens = opens[:len(opens)-1]
// 			stars = stars[:len(stars)-1]
// 		}
// 		return len(opens) == 0
// 	}
// 	return valid(0, len(s), '(', 1) || valid(len(s)-1, -1, ')', -1)
// }

// // time: O(n^2)
// // space: O(n^2)
// func checkValidString(s string) bool {
// 	var n = len(s)
// 	var memo = make([][]*bool, n)
// 	for idx := range memo {
// 		memo[idx] = make([]*bool, n)
// 	}
// 	var dp func(idx int, count int) bool
// 	dp = func(idx int, count int) bool {
// 		if idx == n {
// 			return count == 0
// 		}
// 		if memo[idx][count] != nil {
// 			return *memo[idx][count]
// 		}
// 		var result bool
// 		switch s[idx] {
// 		case '(':
// 			result = dp(idx+1, count+1)
// 		case ')':
// 			result = count > 0 && dp(idx+1, count-1)
// 		default:
// 			result = dp(idx+1, count) || //ignore
// 				dp(idx+1, count+1) || // consider as (
// 				count > 0 && dp(idx+1, count-1) // consider as )
// 		}
// 		memo[idx][count] = &result
// 		return result
// 	}
// 	return dp(0, 0)
// }
