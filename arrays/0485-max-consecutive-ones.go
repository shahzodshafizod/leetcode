package arrays

// https://leetcode.com/problems/max-consecutive-ones/

// Approach: Single pass tracking
// Track current consecutive count and max.
// Time: O(n) - single pass
// Space: O(1) - constant space
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0

	for _, num := range nums {
		if num == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}

	return maxCount
}
