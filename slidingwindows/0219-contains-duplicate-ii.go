package slidingwindows

// https://leetcode.com/problems/contains-duplicate-ii/

// time: O(n)
// space: O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]bool)

	start := 0
	for end, num := range nums {
		if end-start > k {
			delete(window, nums[start])

			start++
		}

		if window[num] {
			return true
		}

		window[num] = true
	}

	return false
}

// // brute force
// // time: O(n^2)
// // space: O(1)
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	var n = len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n && j <= i+k; j++ {
// 			if nums[j] == nums[i] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
