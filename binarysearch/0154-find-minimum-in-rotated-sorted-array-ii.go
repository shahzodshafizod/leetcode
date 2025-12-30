package binarysearch

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

// Approach #2: Optimized - Binary Search with duplicate handling
// Compare mid with right boundary to determine which half to search
// When duplicates cause ambiguity, shrink search space by decrementing right
// N = length of array
// Time: O(log N) average, O(N) worst case (all duplicates)
// Space: O(1) - only use pointers
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] < nums[right] {
			// Minimum is in left half (including mid)
			// Example: [4, 5, 1, 2, 3], mid=1, right=3
			// 1 < 3, so minimum is in left half (could be mid itself)
			right = mid
		} else if nums[mid] > nums[right] {
			// Minimum is in right half (excluding mid)
			// Example: [4, 5, 6, 7, 0, 1, 2], mid=7, right=2
			// 7 > 2, so minimum must be in right half
			left = mid + 1
		} else {
			// nums[mid] == nums[right]
			// Can't determine which half has minimum
			// Example: [2, 2, 2, 0, 1] or [1, 2, 2, 2, 2]
			// Decrement right to shrink search space
			right--
		}
	}

	return nums[left]
}

// // Approach #1: Brute Force - Linear Scan
// // Simply iterate through array to find minimum
// // N = length of array
// // Time: O(N) - scan entire array
// // Space: O(1) - only use min variable
// func findMin(nums []int) int {
// 	minimum := nums[0]
// 	for _, num := range nums {
// 		minimum = min(minimum, num)
// 	}
// 	return minimum
// }
