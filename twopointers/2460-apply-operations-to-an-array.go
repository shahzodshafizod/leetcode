package twopointers

// https://leetcode.com/problems/apply-operations-to-an-array/

// Approach: Brute-Force
// Time: O(n)
// Space: O(1)
func applyOperations(nums []int) []int {
	var n, left = len(nums), 0
	for idx := 0; idx < n; idx++ {
		if nums[idx] == 0 {
			continue
		}
		if idx < n-1 && nums[idx] == nums[idx+1] {
			nums[left] = nums[idx] * 2
			nums[idx+1] = 0
		} else {
			nums[left] = nums[idx]
		}
		left++
	}
	for ; left < n; left++ {
		nums[left] = 0
	}
	return nums
}
