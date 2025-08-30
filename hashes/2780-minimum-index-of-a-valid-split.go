package hashes

// https://leetcode.com/problems/minimum-index-of-a-valid-split/

// Approach #2: Boyer-Moore Majority Voting Algorithm
// Time: O(n)
// Space: O(1)
func minimumIndex(nums []int) int {
	n := len(nums)
	dominant, count := nums[0], 0

	for idx := range n {
		if nums[idx] == dominant {
			count++
		} else {
			count--
		}

		if count == 0 {
			dominant = nums[idx]
			count = 1
		}
	}

	prefix, suffix := 0, 0

	for idx := range n {
		if nums[idx] == dominant {
			suffix++
		}
	}

	for idx := range n {
		if nums[idx] == dominant {
			prefix++
			suffix--
		}

		if prefix*2 > idx+1 && suffix*2 > n-idx-1 {
			return idx
		}
	}

	return -1
}

// // Approach #1: Hash Map
// // Time: O(n)
// // Space: O(n)
// func minimumIndex(nums []int) int {
// 	var n = len(nums)
// 	var dominant = -1
// 	var count = make(map[int]int)
// 	for idx := 0; idx < n; idx++ {
// 		count[nums[idx]]++
// 		if dominant == -1 && count[nums[idx]]*2 > n {
// 			dominant = nums[idx]
// 		}
// 	}
// 	var prefix, suffix = 0, count[dominant]
// 	for idx := 0; idx < n-1; idx++ {
// 		if nums[idx] == dominant {
// 			prefix++
// 			suffix--
// 		}
// 		if prefix*2 > idx+1 && suffix*2 > n-idx-1 {
// 			return idx
// 		}
// 	}
// 	return -1
// }
