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
