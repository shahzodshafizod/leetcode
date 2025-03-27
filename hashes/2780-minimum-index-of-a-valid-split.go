package hashes

// https://leetcode.com/problems/minimum-index-of-a-valid-split/

// Approach #2: Boyer-Moore Majority Voting Algorithm
// Time: O(n)
// Space: O(1)
func minimumIndex(nums []int) int {
	var n = len(nums)
	var dominent, count = nums[0], 0
	for idx := 0; idx < n; idx++ {
		if nums[idx] == dominent {
			count++
		} else {
			count--
		}
		if count == 0 {
			dominent = nums[idx]
			count = 1
		}
	}
	var prefix, suffix = 0, 0
	for idx := 0; idx < n; idx++ {
		if nums[idx] == dominent {
			suffix++
		}
	}
	for idx := 0; idx < n; idx++ {
		if nums[idx] == dominent {
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
// 	var dominent = -1
// 	var count = make(map[int]int)
// 	for idx := 0; idx < n; idx++ {
// 		count[nums[idx]]++
// 		if dominent == -1 && count[nums[idx]]*2 > n {
// 			dominent = nums[idx]
// 		}
// 	}
// 	var prefix, suffix = 0, count[dominent]
// 	for idx := 0; idx < n-1; idx++ {
// 		if nums[idx] == dominent {
// 			prefix++
// 			suffix--
// 		}
// 		if prefix*2 > idx+1 && suffix*2 > n-idx-1 {
// 			return idx
// 		}
// 	}
// 	return -1
// }
