package arrays

// https://leetcode.com/problems/longest-continuous-increasing-subsequence/

// Approach: Optimized - Single pass with counter
// Track current sequence length, reset when not increasing
// Time: O(N) - single pass through array
// Space: O(1) - only storing two counters
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLength := 1
	currentLength := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			currentLength = 1
		}
	}

	return maxLength
}
