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
	getStat := func(a int, b int) int {
		return ((a & 1) << 1) + (b & 1)
	}
	n := len(s)
	calcDifference := func(a byte, b byte) int {
		best := [4]int{0, n, n, n}
		bestB := [4]int{0, n, n, n}
		cnta, cntb := 0, 0
		preva, prevb := 0, 0

		var stat, prevStat, prevDiff int

		maxDiff := -n

		for idx := range n {
			switch s[idx] {
			case a:
				cnta++
			case b:
				cntb++
			default:
			}

			if idx+1 >= k {
				stat = getStat(cnta, cntb) ^ 0b10
				if best[stat] < n && cntb > bestB[stat] {
					// B must have frequency > 0 in the subarray
					maxDiff = max(maxDiff, cnta-cntb-best[stat])
				}
				// slide window to right
				switch s[idx+1-k] {
				case a:
					preva++
				case b:
					prevb++
				default:
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
	vals := [5]byte{'0', '1', '2', '3', '4'}
	maxDiff := -n

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
