package slidingwindows

import "sort"

// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/

// Approach #2: Sliding Window
// Time: O(n log n)
// Space: O(n)
func minOperations(nums []int) int {
	origlen := len(nums)
	unique := make(map[int]struct{})
	for _, num := range nums {
		unique[num] = struct{}{}
	}
	nums = make([]int, 0, len(unique))
	for num := range unique {
		nums = append(nums, num)
	}
	sort.Ints(nums)
	len := len(nums)
	ops := origlen
	end := 0
	var windowSize int
	for start := range nums {
		// target window is: [start; start+length-1]
		for end < len && nums[end] <= nums[start]+origlen-1 {
			end++
		}
		windowSize = end - start
		ops = min(ops, origlen-windowSize)
	}
	return ops
}

// // Approach #1: Binary Search
// // Time: O(n log n)
// // Space: O(n)
// func minOperations(nums []int) int {
// 	var origlen = len(nums)
// 	var unique = make(map[int]struct{})
// 	for _, num := range nums {
// 		unique[num] = struct{}{}
// 	}
// 	nums = make([]int, 0, len(unique))
// 	for num := range unique {
// 		nums = append(nums, num)
// 	}
// 	var len = len(nums)
// 	sort.Ints(nums)
// 	var ops = origlen
// 	var endnum, left, right, mid int
// 	for start := range nums {
// 		endnum = nums[start] + origlen - 1
// 		left, right = 0, len-1
// 		for left <= right {
// 			mid = left + (right-left)/2
// 			if nums[mid] > endnum {
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 		ops = min(ops, origlen-(right-start+1))
// 	}
// 	return ops
// }
