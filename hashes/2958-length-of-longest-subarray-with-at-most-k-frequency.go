package hashes

// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/

// time: O(2n) = O(n)
// space: O(n)
func maxSubarrayLength(nums []int, k int) int {
	count := make(map[int]int) // num:count
	start := 0
	length := 0

	for end := range nums {
		count[nums[end]]++
		for count[nums[end]] > k {
			count[nums[start]]--
			start++
		}

		length = max(length, end-start+1)
	}

	return length
}
