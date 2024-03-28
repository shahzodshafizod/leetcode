package arrays

// https://leetcode.com/problems/subarray-product-less-than-k/

func numSubarrayProductLessThanK(nums []int, k int) int {
	var count, pro = 0, 1
	var start = 0
	for end := range nums {
		pro *= nums[end]
		for pro >= k && start <= end {
			pro /= nums[start]
			start++
		}
		count += end - start + 1
	}
	return count
}
