package arrays

import "strconv"

// https://leetcode.com/problems/summary-ranges/

// Approach: Two pointers for range tracking
// Track start of each range, extend while consecutive.
// Time: O(n) - single pass
// Space: O(1) - excluding result
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := []string{}
	start := 0

	for i := range nums {
		// Check if end of range (last element or gap detected)
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if start == i {
				result = append(result, strconv.Itoa(nums[start]))
			} else {
				result = append(result, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i]))
			}

			start = i + 1
		}
	}

	return result
}
