package arrays

import "sort"

// https://leetcode.com/problems/combination-sum/

// time: O(2^target)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // to skip remaining greater candicates if target is negative
	var combinations = make([][]int, 0)
	var combination = make([]int, 0)
	var length = len(candidates)
	var dfs func(idx int, target int)
	dfs = func(idx int, target int) {
		if target == 0 {
			combinations = append(combinations, append([]int{}, combination...))
			return
		}
		// decision to include candidates[idx]
		target -= candidates[idx]
		if target < 0 {
			return
		}
		combination = append(combination, candidates[idx])
		dfs(idx, target)
		// decision NOT to include candidates[idx]
		combination = combination[:len(combination)-1]
		if idx+1 < length {
			dfs(idx+1, target+candidates[idx])
		}
	}
	dfs(0, target)
	return combinations
}
