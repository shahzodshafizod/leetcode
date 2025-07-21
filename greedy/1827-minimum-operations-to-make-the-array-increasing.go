package greedy

// https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/

// time: O(n)
// space: O(1)
func minOperations(nums []int) int {
	increasings := 0
	for i, n := 1, len(nums); i < n; i++ {
		if nums[i-1] >= nums[i] {
			increasings += 1 + nums[i-1] - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return increasings
}
