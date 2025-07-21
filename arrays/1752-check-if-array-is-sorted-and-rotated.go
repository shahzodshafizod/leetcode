package arrays

// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/

// Approach: Finding Smallest Elements
// Time: Time(n)
// Space: O(1)
func check(nums []int) bool {
	n := len(nums)
	breaks := 0
	for idx := 1; idx < n; idx++ {
		if nums[idx-1] > nums[idx] {
			breaks++
		}
	}
	if nums[n-1] > nums[0] {
		breaks++
	}
	return breaks <= 1
}

// // Approach: Sliding Window
// // Time: O(n)
// // Space: O(1)
// func check(nums []int) bool {
// 	var n = len(nums)
// 	var count = 1
// 	for idx := 2*n - 1; idx > 1; idx-- {
// 		if nums[(idx-1)%n] <= nums[idx%n] {
// 			count++
// 		} else {
// 			count = 1
// 		}
// 		if count == n {
// 			return true
// 		}
// 	}
// 	return n == 1
// }
