package twopointers

// https://leetcode.com/problems/find-all-k-distant-indices-in-an-array/

// Approach: Two Pointers
// Time: O(n)
// Space: O(1)
func findKDistantIndices(nums []int, key int, k int) []int {
	var indices []int

	n, right := len(nums), 0

	var left int

	for idx := 0; idx < n; idx++ {
		if nums[idx] == key {
			left = max(right, idx-k)
			right = min(n-1, idx+k) + 1

			for ; left < right; left++ {
				indices = append(indices, left)
			}
		}
	}

	return indices
}
