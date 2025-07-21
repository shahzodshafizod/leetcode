package slidingwindows

import "slices"

/*
Total number of subarrays (n is len(nums)):
n * (n + 1) / 2
*/

// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

func countSubarrays2962(nums []int, k int) int64 {
	start, max := 0, slices.Max(nums)
	var total int64 = 0
	for _, num := range nums {
		if num == max {
			k--
		}
		for ; k == 0; start++ {
			if nums[start] == max {
				k++
			}
		}
		total += int64(start)
	}
	return total
}
