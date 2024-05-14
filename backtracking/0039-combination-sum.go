package backtracking

import "sort"

// https://leetcode.com/problems/combination-sum/

// time: O(2^target)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // to skip remaining greater candicates if target is negative
	var combinations = make([][]int, 0)
	var length = len(candidates)
	var dfs func(idx int, target int, combination []int)
	dfs = func(idx int, target int, combination []int) {
		if idx < 0 || target < 0 {
			return
		}
		if target == 0 {
			combinations = append(combinations, append([]int{}, combination...))
			return
		}
		// decision to include candidates[idx]
		target -= candidates[idx]
		dfs(idx, target, append(combination, candidates[idx]))
		// decision NOT to include candidates[idx]
		target += candidates[idx]
		dfs(idx-1, target, combination)
	}
	dfs(length-1, target, []int{})
	return combinations
}
