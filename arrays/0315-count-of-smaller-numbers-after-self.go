package arrays

import "slices"

// https://leetcode.com/problems/count-of-smaller-numbers-after-self/

// Approach: Binary Indexed Tree (Fenwick Tree)
// Time: O(N Log N)
// Space: O(N)
func countSmaller(nums []int) []int {
	var min = slices.Min(nums)
	var max = slices.Max(nums)
	var bitree = make([]int, max-min+1+1)
	var update = func(idx int, delta int) {
		idx -= min - 1
		for n := len(bitree); idx < n; idx += idx & -idx {
			bitree[idx] += delta
		}
	}
	var count = func(idx int) int {
		idx -= min
		var count = 0
		for ; idx > 0; idx -= idx & -idx {
			count += bitree[idx]
		}
		return count
	}
	for _, num := range nums {
		update(num, 1)
	}
	var n = len(nums)
	var smallers = make([]int, n)
	for idx, num := range nums {
		update(num, -1)
		smallers[idx] = count(num)
	}
	return smallers
}

// // Approach: Merge Sort
// // Time: O(N Log N)
// // Space: O(N)
// func countSmaller(nums []int) []int {
// 	var n = len(nums)
// 	var indices = make([]int, n)
// 	for idx := range indices {
// 		indices[idx] = idx
// 	}
// 	var smallers = make([]int, n)
// 	var mergeSort func(low int, high int)
// 	mergeSort = func(low int, high int) {
// 		if high-low <= 1 {
// 			return
// 		}
// 		var mid = low + (high-low)/2
// 		mergeSort(low, mid)
// 		mergeSort(mid, high)
// 		var temp = make([]int, high-low)
// 		var i, j, k = low, mid, 0
// 		for i < mid && j < high {
// 			if nums[indices[i]] <= nums[indices[j]] {
// 				temp[k] = indices[i]
// 				smallers[indices[i]] += j - mid
// 				i++
// 			} else {
// 				temp[k] = indices[j]
// 				j++
// 			}
// 			k++
// 		}
// 		for i < mid {
// 			temp[k] = indices[i]
// 			smallers[indices[i]] += j - mid
// 			k++
// 			i++
// 		}
// 		for j < high {
// 			temp[k] = indices[j]
// 			k++
// 			j++
// 		}
// 		copy(indices[low:high], temp)
// 	}
// 	mergeSort(0, n)
// 	return smallers
// }

// // Approach: Brute-Force (Time Limit Exceeded)
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
