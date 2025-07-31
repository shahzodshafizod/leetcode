package greedy

import "slices"

// https://leetcode.com/problems/minimum-increment-to-make-array-unique/

// N = len(nums)
// Len = n + max(nums)
// time: O(N)
// space: O(Len)
func minIncrementForUnique(nums []int) int {
	n := len(nums) + slices.Max(nums)

	count := make([]int, n)
	for _, num := range nums { // O(n)
		count[num]++
	}

	moves := 0

	for num := 0; num < n; num++ { // O(len)
		if count[num] > 1 {
			moves += count[num] - 1
			count[num+1] += count[num] - 1
		}
	}

	return moves
}

// // time: O(n log n)
// // space: O(1)
// func minIncrementForUnique(nums []int) int {
// 	sort.Ints(nums) // O(n log n)
// 	var moves = 0
// 	for i, n := 1, len(nums); i < n; i++ {
// 		if nums[i] <= nums[i-1] {
// 			moves += 1 + nums[i-1] - nums[i]
// 			nums[i] = nums[i-1] + 1
// 		}
// 	}
// 	return moves
// }
