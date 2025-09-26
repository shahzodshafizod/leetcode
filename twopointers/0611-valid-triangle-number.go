package twopointers

import "sort"

// nums[i] <= nums[j] <= nums[k]
// so nums[k] + 'any one' becomes greater than 'another one'

// https://leetcode.com/problems/valid-triangle-number/

// Approach #2: Two Pointers
// Time: O(nn)
// Space: O(1)
func triangleNumber(nums []int) int {
	sort.Ints(nums)

	var total, start, end int
	for i, n := 1, len(nums); i < n; i++ {
		start, end = 0, i-1
		for start < end {
			if nums[start]+nums[end] <= nums[i] {
				start++
			} else {
				total += end - start
				end--
			}
		}
	}

	return total
}

// // Approach #1: Binary Search
// // Time: O(nnlogn)
// // Space: O(1)
// func triangleNumber(nums []int) int {
// 	sort.Ints(nums)

// 	bisectRight := func(combin int, left int, right int) int {
// 		var k int
// 		for left <= right {
// 			k = left + (right-left)/2
// 			if combin <= nums[k] {
// 				right = k - 1
// 			} else {
// 				left = k + 1
// 			}
// 		}

// 		return left
// 	}

// 	var total int

// 	for i, n := 0, len(nums); i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			total += bisectRight(nums[i]+nums[j], j+1, n-1) - j - 1
// 		}
// 	}

// 	return total
// }
