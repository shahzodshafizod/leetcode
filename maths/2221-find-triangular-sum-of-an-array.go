package maths

// https://leetcode.com/problems/find-triangular-sum-of-an-array/

// Approach: Brute-Force (Simulation)
// Time: O(nn)
// Space: O(1)
func triangularSum(nums []int) int {
	for n := len(nums); n > 1; n-- {
		for i := 1; i < n; i++ {
			nums[i-1] = (nums[i-1] + nums[i]) % 10
		}
	}

	return nums[0]
}
