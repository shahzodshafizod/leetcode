package slidingwindows

/*
even - even = even
even - odd = odd
odd - odd = even
odd - even = odd

[e]e - [o]e = oe
[e]o - [o]o = oe
[o]e - [e]e = oe
[o]o - [e]o = oe
it means the second remains the same, but the first changes
*/

// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-ii/

// Approach #2: Prefix Sum + Sliding Window
// Time: O(n + mm), m=5
// Space: O(n)
func maxDifference(s string, k int) int {
	var getStat = func(a int, b int) int {
		return ((a & 1) << 1) + (b & 1)
	}
	var n = len(s)
	var calcDifference = func(a byte, b byte) int {
		var best = [4]int{0, n, n, n}
		var bestB = [4]int{0, n, n, n}
		var cnta, cntb = 0, 0
		var preva, prevb = 0, 0
		var stat, prevStat, prevDiff int
		var maxDiff = -n
		for idx := 0; idx < n; idx++ {
			if s[idx] == a {
				cnta++
			} else if s[idx] == b {
				cntb++
			}
			if idx+1 >= k {
				stat = getStat(cnta, cntb) ^ 0b10
				if best[stat] < n && cntb > bestB[stat] {
					// B must have frequency > 0 in the subarray
					maxDiff = max(maxDiff, cnta-cntb-best[stat])
				}
				// slide window to right
				if s[idx+1-k] == a {
					preva++
				} else if s[idx+1-k] == b {
					prevb++
				}
				prevStat = getStat(preva, prevb)
				prevDiff = preva - prevb
				if prevDiff < best[prevStat] {
					best[prevStat] = prevDiff
					bestB[prevStat] = prevb
				}
			}
		}
		return maxDiff
	}
	var vals = [5]byte{'0', '1', '2', '3', '4'}
	var maxDiff = -n
	for _, a := range vals {
		for _, b := range vals {
			if a != b {
				maxDiff = max(maxDiff, calcDifference(a, b))
			}
		}
	}
	return maxDiff
}

// // Approach #1: Sliding Window
// // Time: O(nn)
// // Space: O(1)
// func maxDifference(s string, k int) int {
// 	var n = len(s)
// 	var diff = -n
// 	var odd, even, idx, cnt int
// 	for size := k; size <= n; size++ {
// 		var count [5]int
// 		for idx = size - 2; idx >= 0; idx-- {
// 			count[int(s[idx]-'0')]++
// 		}
// 		for idx = size - 1; idx < n; idx++ {
// 			count[int(s[idx]-'0')]++
// 			odd, even = 0, n
// 			for _, cnt = range count {
// 				if cnt&1 == 1 {
// 					odd = max(odd, cnt)
// 				} else if cnt > 0 {
// 					even = min(even, cnt)
// 				}
// 			}
// 			diff = max(diff, odd-even)
// 			count[int(s[idx-size+1]-'0')]--
// 		}
// 	}
// 	return diff
// }
