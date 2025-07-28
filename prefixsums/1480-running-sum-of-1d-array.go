package prefixsums

// https://leetcode.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	for i, n := 1, len(nums); i < n; i++ {
		nums[i] += nums[i-1]
	}

	return nums
}
