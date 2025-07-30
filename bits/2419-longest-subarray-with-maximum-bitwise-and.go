package bits

// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/

// Approach #2: One-Pass
// Time: O(n)
// Space: O(1)
func longestSubarray(nums []int) int {
	// 1. if n < cur_max, n & cur_max < cur_max
	// 2. if n == cur_max, n & cur_max = cur_max
	// 3. if n > cur_max, n & cur_max < cur_max
	curmax := 0
	length, maxlen := 0, 0

	for _, num := range nums {
		if num > curmax {
			curmax = num
			length = 1
			maxlen = 0
		} else if num < curmax {
			length = 0
		} else {
			length++
		}

		maxlen = max(maxlen, length)
	}

	return maxlen
}

// // Approach #1: Two-Pass
// // Time: O(n)
// // Space: O(1)
// func longestSubarray(nums []int) int {
// 	target := slices.Max(nums)
// 	length, maxlen := 0, 0

// 	for _, num := range nums {
// 		if num == target {
// 			length++
// 		} else {
// 			length = 0
// 		}

// 		maxlen = max(maxlen, length)
// 	}

// 	return maxlen
// }
