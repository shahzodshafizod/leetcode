package twopointers

// https://leetcode.com/problems/count-binary-substrings/

// Approach: Optimized - Group counting with tracking
// Count consecutive groups, then calculate valid substrings between adjacent groups
// For groups of size m and n, can form min(m, n) valid substrings
// Example: "000111" -> groups [3, 3] -> min(3, 3) = 3 substrings
// N = length of string
// Time: O(N) - single pass through string
// Space: O(1) - only track previous and current group counts
func countBinarySubstrings(s string) int {
	prevCount := 0
	currCount := 1
	result := 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			// Same character, extend current group
			currCount++
		} else {
			// Different character, start new group
			// Add min(prev, curr) valid substrings
			result += min(prevCount, currCount)
			prevCount = currCount
			currCount = 1
		}
	}

	// Don't forget the last group
	result += min(prevCount, currCount)

	return result
}
