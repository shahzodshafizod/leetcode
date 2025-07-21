package arrays

import "sort"

// https://leetcode.com/problems/sort-array-by-increasing-frequency/

// Time: O(N Log N)
// Space: O(N)
func frequencySort(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}

// // Time: O(N Log N)
// // Space: O(N)
// func frequencySort(nums []int) []int {
// 	var freq [201]int
// 	for _, num := range nums {
// 		freq[num+100]++
// 	}
// 	sort.Slice(nums, func(i, j int) bool {
// 		if freq[nums[i]+100] == freq[nums[j]+100] {
// 			return nums[i] > nums[j]
// 		}
// 		return freq[nums[i]+100] < freq[nums[j]+100]
// 	})
// 	return nums
// }
