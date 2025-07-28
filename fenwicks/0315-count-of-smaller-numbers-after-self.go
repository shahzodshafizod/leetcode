package fenwicks

import "slices"

// https://leetcode.com/problems/count-of-smaller-numbers-after-self/

// Approach: Binary Indexed Tree (Fenwick Tree)
// Time: O(N Log N)
// Space: O(N)
func countSmaller(nums []int) []int {
	min := slices.Min(nums)
	bitree := make([]int, slices.Max(nums)-min+2) // 1-indexed
	update := func(idx int, delta int) {
		idx -= min - 1
		for n := len(bitree); idx < n; idx += idx & -idx {
			bitree[idx] += delta
		}
	}
	count := func(idx int) int {
		idx -= min
		count := 0

		for ; idx > 0; idx -= idx & -idx {
			count += bitree[idx]
		}

		return count
	}

	for _, num := range nums {
		update(num, 1)
	}

	n := len(nums)
	smallers := make([]int, n)

	for idx, num := range nums {
		update(num, -1)
		smallers[idx] = count(num)
	}

	return smallers
}

// // Approach: Merge Sort
// // Time: O(NLogN)
// // Space: O(N)
// func countSmaller(nums []int) []int {
// 	var n = len(nums)
// 	var indices = make([]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		indices[idx] = idx
// 	}
// 	var counts = make([]int, n)
// 	var merge func(left int, right int)
// 	merge = func(left int, right int) {
// 		if right-left <= 1 {
// 			return
// 		}
// 		var mid = left + (right-left)/2
// 		merge(left, mid)
// 		merge(mid, right)
// 		var count = 0
// 		var merged = make([]int, right-left)
// 		var left1, left2 = left, mid
// 		for idx := 0; left1 < mid || left2 < right; idx++ {
// 			if left1 == mid || left2 < right && nums[indices[left2]] < nums[indices[left1]] {
// 				merged[idx] = indices[left2]
// 				left2++
// 				count++
// 			} else {
// 				merged[idx] = indices[left1]
// 				counts[indices[left1]] += count
// 				left1++
// 			}
// 		}
// 		copy(indices[left:right], merged)
// 	}
// 	merge(0, n)
// 	return counts
// }

// // Approach: Brute Force (Time Limit Exceeded)
// // Time: O(N^2)
// // Space: O(1)
// func countSmaller(nums []int) []int {
// 	var n = len(nums)
// 	var rSmallers = make([]int, n)
// 	for idx := n - 2; idx >= 0; idx-- {
// 		for jdx := idx + 1; jdx < n; jdx++ {
// 			if nums[jdx] < nums[idx] {
// 				rSmallers[idx]++
// 			}
// 		}
// 	}
// 	return rSmallers
// }
