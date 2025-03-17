package hashes

import "sort"

// https://leetcode.com/problems/divide-array-into-equal-pairs/

// Approach #2: Sorting
// Time: O(nlogn)
// Space: O(1)
func divideArray(nums []int) bool {
	sort.Ints(nums)
	var prev, count = -1, 0
	for _, num := range nums {
		if num != prev && count&1 != 0 {
			break
		}
		count++
		prev = num
	}
	return count&1 == 0
}

// // Approach #1: Hashes, Counting
// // Time: O(n)
// // Space: O(n)
// func divideArray(nums []int) bool {
// 	var count = make(map[int]int)
// 	for _, num := range nums {
// 		count[num]++
// 	}
// 	for _, cnt := range count {
// 		if cnt&1 != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
