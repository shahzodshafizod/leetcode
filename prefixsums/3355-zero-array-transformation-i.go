package prefixsums

// https://leetcode.com/problems/zero-array-transformation-i/

// Approach: Line Sweep
// Time: O(n+m), n=len(nums), m=len(queries)
// Space: O(n)
func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	line := make([]int, n+1)
	for idx := range queries {
		line[queries[idx][0]]++
		line[queries[idx][1]+1]--
	}
	presum := 0
	for idx := 0; idx < n; idx++ {
		presum += line[idx]
		if nums[idx] > presum {
			return false
		}
	}
	return true
}
