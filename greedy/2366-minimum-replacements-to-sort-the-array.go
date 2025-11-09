package greedy

// https://leetcode.com/problems/minimum-replacements-to-sort-the-array/

// Approach: Greedy
// Time: O(n)
// Space: O(1)
func minimumReplacement(nums []int) int64 {
	var ops int64

	n := len(nums)

	var parts int

	height := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		parts = (nums[i] + height - 1) / height
		ops += int64(parts) - 1
		height = nums[i] / parts
	}

	return ops
}
