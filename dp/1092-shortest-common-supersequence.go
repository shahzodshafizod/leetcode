package dp

import "strings"

// https://leetcode.com/problems/shortest-common-supersequence/

// Approach #4: Bottom-Up Dynamic Programming (Memory Optimized)
// Time: O(n⋅m)
// Space: O(n⋅m)
func shortestCommonSupersequence(str1 string, str2 string) string {
	n1, n2 := len(str1), len(str2)
	dp := make([][]int, n1+1)

	for idx := range dp {
		dp[idx] = make([]int, n2+1)
	}

	var i1, i2 int
	for i1 = n1 - 1; i1 >= 0; i1-- {
		for i2 = n2 - 1; i2 >= 0; i2-- {
			if str1[i1] == str2[i2] {
				dp[i1][i2] = 1 + dp[i1+1][i2+1]
			} else {
				dp[i1][i2] = max(dp[i1+1][i2], dp[i1][i2+1])
			}
		}
	}

	var (
		csc strings.Builder
		err error
	)

	i1, i2 = 0, 0
	for i1 < n1 && i2 < n2 {
		if str1[i1] == str2[i2] {
			err = csc.WriteByte(str1[i1])

			i1++
			i2++
		} else if dp[i1][i2+1] < dp[i1+1][i2] {
			err = csc.WriteByte(str1[i1])

			i1++
		} else {
			err = csc.WriteByte(str2[i2])

			i2++
		}

		_ = err
	}

	_, err = csc.WriteString(str1[i1:])
	_ = err
	_, err = csc.WriteString(str2[i2:])
	_ = err

	return csc.String()
}

// // Approach #3: Bottom-Up Dynamic Programming (Time Optimized)
// // Time: O(n⋅m)
// // Space: O(n⋅m)
// func shortestCommonSupersequence(str1 string, str2 string) string {
// 	var n1, n2 = len(str1), len(str2)
// 	var curr = make([]string, n2+1)
// 	for i2 := n2 - 1; i2 >= 0; i2-- {
// 		curr[i2] = string(str2[i2]) + curr[i2+1]
// 	}
// 	for i1 := n1 - 1; i1 >= 0; i1-- {
// 		var next = curr
// 		curr = make([]string, n2+1)
// 		curr[n2] = str1[i1:]
// 		for i2 := n2 - 1; i2 >= 0; i2-- {
// 			if str1[i1] == str2[i2] {
// 				curr[i2] = string(str1[i1]) + next[i2+1]
// 			} else {
// 				substr1 := string(str1[i1]) + next[i2]
// 				substr2 := string(str2[i2]) + curr[i2+1]
// 				if len(substr1) < len(substr2) {
// 					curr[i2] = substr1
// 				} else {
// 					curr[i2] = substr2
// 				}
// 			}
// 		}
// 	}
// 	return curr[0]
// }

// // Approach #2: Top-Down Dynamic Programming (Memory Limit Exceeded)
// // Time: O(n⋅m⋅(n+m)), n=len(str1), m=len(str2)
// // Space: O(n⋅m⋅(n+m))
// func shortestCommonSupersequence(str1 string, str2 string) string {
// 	var n1, n2 = len(str1), len(str2)
// 	var memo = make([][]*string, n1)
// 	for idx := range memo {
// 		memo[idx] = make([]*string, n2)
// 	}
// 	var backtrack func(i1 int, i2 int) string
// 	backtrack = func(i1 int, i2 int) string {
// 		if i1 == n1 {
// 			return str2[i2:]
// 		}
// 		if i2 == n2 {
// 			return str1[i1:]
// 		}
// 		if memo[i1][i2] != nil {
// 			return *memo[i1][i2]
// 		}
// 		if str1[i1] == str2[i2] {
// 			csc := string(str1[i1]) + backtrack(i1+1, i2+1)
// 			memo[i1][i2] = &csc
// 			return csc
// 		}
// 		substr1 := string(str1[i1]) + backtrack(i1+1, i2)
// 		substr2 := string(str2[i2]) + backtrack(i1, i2+1)
// 		if len(substr1) < len(substr2) {
// 			memo[i1][i2] = &substr1
// 		} else {
// 			memo[i1][i2] = &substr2
// 		}
// 		return *memo[i1][i2]
// 	}
// 	return backtrack(0, 0)
// }

// // Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
// // Time: O(2^(n+m)⋅(n+m)), n=len(str1), m=len(str2)
// // Space: O((n+m)^2) for recursion stack
// func shortestCommonSupersequence(str1 string, str2 string) string {
// 	var n1, n2 = len(str1), len(str2)
// 	var backtrack func(i1 int, i2 int) string
// 	backtrack = func(i1 int, i2 int) string {
// 		if i1 == n1 {
// 			return str2[i2:]
// 		}
// 		if i2 == n2 {
// 			return str1[i1:]
// 		}
// 		if str1[i1] == str2[i2] {
// 			return string(str1[i1]) + backtrack(i1+1, i2+1)
// 		}
// 		substr1 := string(str1[i1]) + backtrack(i1+1, i2)
// 		substr2 := string(str2[i2]) + backtrack(i1, i2+1)
// 		if len(substr1) < len(substr2) {
// 			return substr1
// 		}
// 		return substr2
// 	}
// 	return backtrack(0, 0)
// }
