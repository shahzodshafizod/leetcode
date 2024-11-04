package strings

// https://leetcode.com/problems/rotate-string/

// Approach #3: KMP (Knuth-Morris-Pratt) Algorithm
// Time: O(n)
// Space: O(n)
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	// 1. LPS setup: O(2xn)
	var n = len(goal)
	var lps = make([]int, n) // longest prefix-suffix
	var prevLps = 0
	for gid := 1; gid < n; gid++ {
		for prevLps > 0 && goal[prevLps] != goal[gid] {
			prevLps = lps[prevLps-1]
		}
		if goal[gid] == goal[prevLps] {
			prevLps++
		}
		lps[gid] = prevLps
	}
	// 2. KMP: O(2xn)
	var gid = 0
	for _, c := range s + s {
		for gid > 0 && goal[gid] != byte(c) {
			gid = lps[gid-1]
		}
		if goal[gid] == byte(c) {
			gid++
		}
		if gid == n {
			return true
		}
	}
	return false
}

// // Approach #2: Concatenation
// // Time: O(n)
// // Space: O(n)
// func rotateString(s string, goal string) bool {
// 	return len(s) == len(goal) && strings.Contains(s+s, goal)
// }

// // Approach #1: Brute-Force
// // Time: O(nn)
// // Space: O(1)
// func rotateString(s string, goal string) bool {
// 	if len(s) != len(goal) {
// 		return false
// 	}
// 	var n = len(s)
// 	var sid, gid int
// 	for start := 0; start < n; start++ {
// 		sid, gid = start, 0
// 		for ; gid < n; gid++ {
// 			if s[sid] != goal[gid] {
// 				break
// 			}
// 			sid = (sid + 1) % n
// 		}
// 		if gid == n {
// 			return true
// 		}
// 	}
// 	return false
// }
