package hashes

// https://leetcode.com/problems/contains-duplicate/

// time: O(n)
// space: O(n)
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// // Approach: sort and check every two adjacent neighbors
// // time: O(n log n)
// // space: O(1)
// func containsDuplicate(nums []int) bool {
// 	sort.Ints(nums)
// 	for idx := len(nums) - 2; idx >= 0; idx-- {
// 		if nums[idx] == nums[idx+1] {
// 			return true
// 		}
// 	}
// 	return false
// }
