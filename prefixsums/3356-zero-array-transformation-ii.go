package prefixsums

// https://leetcode.com/problems/zero-array-transformation-ii/

// Approach: Line Sweep
// Time: O(n+m), n=len(nums), m=len(queries)
// Space: O(n)
func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	line := make([]int, n+1)
	k, qlen := 0, len(queries)
	idx, presum := 0, 0
	var l, r, val int
	for idx < n {
		if nums[idx] <= presum+line[idx] {
			presum += line[idx]
			idx++
		} else if k < qlen {
			l, r, val = queries[k][0], queries[k][1], queries[k][2]
			if r >= idx {
				if l <= idx {
					presum += val
				} else {
					line[l] += val
				}
				line[r+1] -= val
			}
			k++
		} else {
			break
		}
	}
	if idx == n {
		return k
	}
	return -1
}
