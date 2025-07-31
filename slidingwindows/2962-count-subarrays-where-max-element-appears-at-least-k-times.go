package slidingwindows

import "slices"

/*
Total number of subarrays (n is len(nums)):
n * (n + 1) / 2
*/

// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

func countSubarrays2962(nums []int, k int) int64 {
	start, maximum := 0, slices.Max(nums)

	var total int64

	for _, num := range nums {
		if num == maximum {
			k--
		}

		for ; k == 0; start++ {
			if nums[start] == maximum {
				k++
			}
		}

		total += int64(start)
	}

	return total
}
