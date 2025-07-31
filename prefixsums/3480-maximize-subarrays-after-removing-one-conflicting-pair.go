package prefixsums

import "slices"

// https://leetcode.com/problems/maximize-subarrays-after-removing-one-conflicting-pair/

// Time: O(n)
// Space: O(n)
func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	right := make([][]int, n+1)

	var mx, mn int
	for _, pair := range conflictingPairs {
		mx, mn = max(pair[0], pair[1]), min(pair[0], pair[1])
		right[mx] = append(right[mx], mn)
	}

	var res int64

	left := [2]int{0, 0}
	ar := make([]int64, n+1) // after removal

	for r := 1; r <= n; r++ {
		for _, l := range right[r] {
			// left = max(left, [l, left[0]], [left[0], l])
			if l > left[0] {
				left[1] = left[0]
				left[0] = l
			} else if l > left[1] {
				left[1] = l
			}
		}

		res += int64(r - left[0])
		ar[left[0]] += int64(left[0] - left[1])
	}

	return res + slices.Max(ar)
}
