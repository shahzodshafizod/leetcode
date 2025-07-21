package arrays

// https://leetcode.com/problems/maximum-ascending-subarray-sum/

func maxAscendingSum(nums []int) int {
	sum, maxSum := nums[0], nums[0]
	for idx, n := 1, len(nums); idx < n; idx++ {
		if nums[idx-1] < nums[idx] {
			sum += nums[idx]
		} else {
			sum = nums[idx]
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
