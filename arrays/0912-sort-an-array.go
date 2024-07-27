package arrays

import "math"

// https://leetcode.com/problems/sort-an-array/

// Approach: Counting Sort
// Time: O(3*N) = O(N)
// Space: O(N)
func sortArray(nums []int) []int {
	var mini, maxi = math.MaxInt, math.MinInt
	for _, num := range nums {
		mini = min(mini, num)
		maxi = max(maxi, num)
	}
	var count = make([]int, maxi-mini+1)
	for _, num := range nums {
		count[num-mini]++
	}
	var idx = 0
	for num, count := range count {
		num += mini
		for count > 0 {
			count--
			nums[idx] = num
			idx++
		}
	}
	return nums
}

// // Approach: Heap Sort
// // Time: O(N Log N)
// // Space: O(1)
// func sortArray(nums []int) []int {
// 	var maxHeap = design.NewPQSort(
// 		len(nums),
// 		func(i, j int) bool { return nums[i] < nums[j] },
// 		func(i, j int) { nums[i], nums[j] = nums[j], nums[i] },
// 	)
// 	maxHeap.Sort()
// 	return nums
// }
