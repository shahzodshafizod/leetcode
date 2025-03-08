package slidingwindows

import "math"

// https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/

// Approach: Sliding Window
// Time: O(n)
// Space: O(1)
func minimumRecolors(blocks string, k int) int {
	var res = math.MaxInt
	ops, k := 0, k-1
	for idx := range blocks {
		if blocks[idx] == 'W' {
			ops++
		}
		if idx-k >= 0 {
			res = min(res, ops)
			if blocks[idx-k] == 'W' {
				ops--
			}
		}
	}
	return res
}

// // Approach: Backtracking
// // Time: O(2^k)
// // Space: O(2^k)
// func minimumRecolors(blocks string, k int) int {
// 	var n = len(blocks)
// 	var backtrack func(idx int, length int) int
// 	backtrack = func(idx int, length int) int {
// 		if length == k {
// 			return 0
// 		}
// 		if idx == n || length+(n-idx-1) < k {
// 			return 100
// 		}
// 		if blocks[idx] == 'B' {
// 			return backtrack(idx+1, length+1)
// 		}
// 		return min(
// 			1+backtrack(idx+1, length+1),
// 			backtrack(idx+1, 0),
// 		)
// 	}
// 	return backtrack(0, 0)
// }
