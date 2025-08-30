package twopointers

// https://leetcode.com/problems/partition-array-according-to-given-pivot/

// Approach: In-place
// Time: O(n)
// Space: O(n)
func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	left := 0

	for idx := range n {
		if nums[idx] < pivot {
			nums[left] = nums[idx]
			left++
		} else if nums[idx] > pivot {
			nums = append(nums, nums[idx])
		}
	}

	for ; left < n; left++ {
		nums[left] = pivot
	}

	greater := len(nums) - n
	copy(nums[n-greater:], nums[n:])
	nums = nums[:n]

	return nums
}

// // Approach: Additional space
// // Time: O(n)
// // Space: O(n)
// func pivotArray(nums []int, pivot int) []int {
// 	var n = len(nums)
// 	var ans = make([]int, n)
// 	var left, right = 0, n - 1
// 	for idx := 0; idx < n; idx++ {
// 		if nums[idx] < pivot {
// 			ans[left] = nums[idx]
// 			left++
// 		}
// 		if nums[n-idx-1] > pivot {
// 			ans[right] = nums[n-idx-1]
// 			right--
// 		}
// 	}
// 	for ; left <= right; left++ {
// 		ans[left] = pivot
// 	}
// 	return ans
// }
