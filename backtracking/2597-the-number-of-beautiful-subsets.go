package backtracking

import "sort"

// https://leetcode.com/problems/the-number-of-beautiful-subsets/

// time: O(n)
// space: O(n)
func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	var count = make(map[int][]int)
	var mod int
	for _, num := range nums {
		mod = num % k
		count[mod] = append(count[mod], num)
	}
	var subsets = 1
	for _, group := range count {
		var subset = 1
		var dict = make(map[int]int)
		for _, num := range group {
			cnt := subset - dict[num-k]
			subset += cnt
			dict[num] += cnt
		}
		subsets *= subset
	}
	return subsets - 1
}

// // time: O(2^n)
// // space: O(n)
// func beautifulSubsets(nums []int, k int) int {
// 	var dfs func(i int, count map[int]int) int
// 	dfs = func(i int, count map[int]int) int {
// 		if i == len(nums) {
// 			return 1
// 		}
// 		// decision NOT to include nums[i]
// 		var subsets = dfs(i+1, count)
// 		// decision to include nums[i]
// 		if count[nums[i]+k] == 0 && count[nums[i]-k] == 0 {
// 			count[nums[i]]++
// 			subsets += dfs(i+1, count)
// 			count[nums[i]]--
// 		}
// 		return subsets
// 	}
// 	return dfs(0, make(map[int]int)) - 1
// }
