package maths

import "strings"

// https://leetcode.com/problems/kth-smallest-instructions/

// Approach: Iterative
// Time: O((row+col)^2)
// Space: O(1)
func kthSmallestPath(destination []int, k int) string {
	var dp [30][30]int
	// Credit to: http://ptaskbook.com/en/tasks/recur.php (Recur6)
	var comb func(n int, k int) int
	comb = func(n int, k int) int {
		if n == k || k == 0 {
			return 1
		}
		if dp[n][k] != 0 {
			return dp[n][k]
		}
		dp[n][k] = comb(n-1, k) + comb(n-1, k-1)
		return dp[n][k]
	}
	row, col := destination[0], destination[1]
	path := ""
	for row != 0 && col != 0 {
		c := comb(row+col-1, row)
		if k <= c {
			path += "H"
			col--
		} else {
			k -= c
			path += "V"
			row--
		}
	}
	path += strings.Repeat("H", col) + strings.Repeat("V", row)
	return path
}

// // Approach: Recursive
// // Time: O((row+col)^2)
// // Space: O(row+col)
// func kthSmallestPath(destination []int, k int) string {
// 	var dp [30][30]int
// // Credit to: http://ptaskbook.com/en/tasks/recur.php (Recur6)
// 	var comb func(n int, k int) int
// 	comb = func(n int, k int) int {
// 		if n == k || k == 0 {
// 			return 1
// 		}
// 		if dp[n][k] != 0 {
// 			return dp[n][k]
// 		}
// 		dp[n][k] = comb(n-1, k) + comb(n-1, k-1)
// 		return dp[n][k]
// 	}
// 	var recur func(row int, col int, k int) string
// 	recur = func(row int, col int, k int) string {
// 		if row == 0 {
// 			return strings.Repeat("H", col)
// 		}
// 		if col == 0 {
// 			return strings.Repeat("V", row)
// 		}
// 		c := comb(row+col-1, row)
// 		if k <= c {
// 			return "H" + recur(row, col-1, k)
// 		}
// 		return "V" + recur(row-1, col, k-c)
// 	}
// 	return recur(destination[0], destination[1], k)
// }

// // Approach: Brute-Force DFS (Time+Memory Limit Exceeded Exceptions)
// func kthSmallestPath(destination []int, k int) string {
// 	var rdest, cdest = destination[0], destination[1]
// 	var dfs func(row int, col int, path string) string
// 	dfs = func(row int, col int, path string) string {
// 		if row > rdest || col > cdest {
// 			return ""
// 		}
// 		if row == rdest && col == cdest {
// 			k--
// 			if k == 0 {
// 				return path
// 			}
// 			return ""
// 		}
// 		for _, next := range []struct {
// 			row  int
// 			col  int
// 			path string
// 		}{
// 			{row, col + 1, path + "H"}, {row + 1, col, path + "V"},
// 		} {
// 			result := dfs(next.row, next.col, next.path)
// 			if result != "" {
// 				return result
// 			}
// 		}
// 		return ""
// 	}
// 	return dfs(0, 0, "")
// }
