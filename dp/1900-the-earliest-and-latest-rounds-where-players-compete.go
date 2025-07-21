package dp

import "math"

// https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/

// Approach #1: Counting + Bitmask + Dynamic Programming
// Time: O(n^4)
// Space: O(n)
func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	f, s := firstPlayer-1, secondPlayer-1
	earliest, latest := math.MaxInt, 0
	var seen [28][28][28]bool
	var ldelta, mdelta, rdelta int
	var dfs func(left int, right int, mask int, round int, lcnt int, mcnt int, rcnt int)
	dfs = func(left int, right int, mask int, round int, lcnt int, mcnt int, rcnt int) {
		if left >= right {
			// We're done with the current round. Go to the next one
			dfs(0, n-1, mask, round+1, lcnt, mcnt, rcnt)
		} else if mask&(1<<left) == 0 {
			// Skip the defeated left player
			dfs(left+1, right, mask, round, lcnt, mcnt, rcnt)
		} else if mask&(1<<right) == 0 {
			// Skip the defeated right player
			dfs(left, right-1, mask, round, lcnt, mcnt, rcnt)
		} else if left == f && right == s {
			// Stop, f & s are fighting against each other
			earliest = min(earliest, round)
			latest = max(latest, round)
		} else if !seen[lcnt][mcnt][rcnt] {
			seen[lcnt][mcnt][rcnt] = true
			if left != f && left != s {
				// Let left LOSE and right WIN
				ldelta, mdelta, rdelta = 0, 0, 0
				if left < f {
					ldelta = 1
				}
				if f < left && left < s {
					mdelta = 1
				}
				if s < left {
					rdelta = 1
				}
				dfs(left+1, right-1, mask^(1<<left), round, lcnt-ldelta, mcnt-mdelta, rcnt-rdelta)
			}
			if right != f && right != s {
				// Let right LOSE and left WIN
				ldelta, mdelta, rdelta = 0, 0, 0
				if right < f {
					ldelta = 1
				}
				if f < right && right < s {
					mdelta = 1
				}
				if s < right {
					rdelta = 1
				}
				dfs(left+1, right-1, mask^(1<<right), round, lcnt-ldelta, mcnt-mdelta, rcnt-rdelta)
			}
		}
	}
	dfs(0, n-1, (1<<n)-1, 1, f, s-f-1, n-1-s)
	return []int{earliest, latest}
}

// // Approach #1: Bitmask + Dynamic Programming
// // Time: O(2^n)
// // Space: O(2^n)
// func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
// 	var f, s = firstPlayer-1, secondPlayer-1
// 	var earliest, latest = math.MaxInt, 0
// 	var dfs func(left int, right int, mask int, round int)
// 	dfs = func(left int, right int, mask int, round int) {
// 		if left >= right {
// 			// We're done with the current round. Go to the next one
// 			dfs(0, n-1, mask, round+1)
// 		} else if mask&(1<<left) == 0 {
// 			// Skip the defeated left player
// 			dfs(left+1, right, mask, round)
// 		} else if mask&(1<<right) == 0 {
// 			// Skip the defeated right player
// 			dfs(left, right-1, mask, round)
// 		} else if left == f && right == s {
// 			// Stop, f & s are fighting against each other
// 			earliest = min(earliest, round)
// 			latest = max(latest, round)
// 		} else {
// 			if left != f && left != s {
// 				// Let left LOSE and right WIN
// 				dfs(left+1, right-1, mask^(1<<left), round)
// 			}
// 			if right != f && right != s {
// 				// Let right LOSE and left WIN
// 				dfs(left+1, right-1, mask^(1<<right), round)
// 			}
// 		}
// 	}
// 	dfs(0, n-1, (1<<n)-1, 1)
// 	return []int{earliest, latest}
// }
