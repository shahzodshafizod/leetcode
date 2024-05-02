package hashes

import "sort"

// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/

// time: O(n log n)
// space: O(1)
func findMaxK(nums []int) int {
	sort.Ints(nums)
	for left, right := 0, len(nums)-1; left < right; {
		if nums[left] > 0 || nums[right] < 0 {
			break
		}
		if -nums[left] == nums[right] {
			return nums[right]
		}
		if -nums[left] > nums[right] {
			left++
		} else {
			right--
		}
	}
	return -1
}

// // time: O(n)
// // space: O(m)
// func findMaxK(nums []int) int {
// 	var bitset [1001]int16
// 	for _, num := range nums {
// 		if num < 0 {
// 			bitset[-num] |= 1
// 		} else {
// 			bitset[num] |= 2
// 		}
// 	}
// 	for idx := 1000; idx > 0; idx-- {
// 		if bitset[idx] == 3 {
// 			return idx
// 		}
// 	}
// 	return -1
// }

// // time: O(n)
// // space: O(n)
// func findMaxK(nums []int) int {
// 	var seen = make(map[int]bool)
// 	var k = -1
// 	var pos int
// 	for _, num := range nums {
// 		pos = int(math.Abs(float64(num)))
// 		if seen[-num] && pos > k {
// 			k = pos
// 		}
// 		seen[num] = true
// 	}
// 	return k
// }
