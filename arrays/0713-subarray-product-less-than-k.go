package arrays

// https://leetcode.com/problems/subarray-product-less-than-k/

func numSubarrayProductLessThanK(nums []int, k int) int {
	var count, pro = 0, 1
	var start = 0
	for end, num := range nums {
		pro *= num
		for pro >= k && start < end {
			pro /= nums[start]
			start++
		}
		if pro < k {
			count += end - start + 1
		}
	}
	return count
}
