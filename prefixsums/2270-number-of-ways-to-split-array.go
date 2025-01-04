package prefixsums

// https://leetcode.com/problems/number-of-ways-to-split-array/

func waysToSplitArray(nums []int) int {
	var right = 0
	for _, num := range nums {
		right += num
	}
	var left = 0
	var count = 0
	for idx, n := 0, len(nums); idx < n-1; idx++ {
		left += nums[idx]
		right -= nums[idx]
		if left >= right {
			count++
		}
	}
	return count
}
