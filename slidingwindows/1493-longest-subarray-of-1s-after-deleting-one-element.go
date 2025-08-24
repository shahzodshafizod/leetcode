package slidingwindows

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

// Approach #2: Sliding Window
// Time: O(n)
// Space: O(1)
func longestSubarray1493(nums []int) int {
	size, start, tries := 0, 0, 1

	for end := range len(nums) {
		if nums[end] == 0 {
			tries--
		}

		for ; tries < 0; start++ {
			if nums[start] == 0 {
				tries++
			}
		}

		size = max(size, end-start)
	}

	return size
}

// // Approach #1: Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func longestSubarray1493(nums []int) int {
// 	n := len(nums)

// 	memo := make([][][2]int, n)
// 	for idx := range memo {
// 		memo[idx] = make([][2]int, n)
// 	}

// 	var dp func(idx int, size int, tries int) int

// 	dp = func(idx, size, tries int) int {
// 		if idx == n {
// 			return size - tries
// 		}

// 		if memo[idx][size][tries] != 0 {
// 			return memo[idx][size][tries]
// 		}

// 		res := size
// 		if nums[idx] == 1 {
// 			res = max(res, dp(idx+1, size+1, tries))
// 		} else if tries > 0 {
// 			res = max(res, dp(idx+1, size, 0), dp(idx+1, 0, 1))
// 		}

// 		memo[idx][size][tries] = res

// 		return res
// 	}

// 	return dp(0, 0, 1)
// }
