package slidingwindows

// https://leetcode.com/problems/longest-nice-subarray/

// Approach: Sliding Window + Bit Manipulation
// Time: O(n)
// Space: O(1)
func longestNiceSubarray(nums []int) int {
	var mask, length = 0, 0
	var left = 0
	for right := range nums {
		for mask&nums[right] != 0 {
			mask ^= nums[left]
			left++
		}
		mask |= nums[right]
		length = max(length, right-left+1)
	}
	return length
}
