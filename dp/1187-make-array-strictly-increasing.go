package dp

// https://leetcode.com/problems/make-array-strictly-increasing/

import (
	"sort"
)

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(n * m * log(m)) where n = len(arr1), m = len(arr2)
// Space: O(m) for the dp map
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	set := make(map[int]struct{})
	for _, num := range arr2 {
		set[num] = struct{}{}
	}

	arr2 = make([]int, 0, len(set))
	for num := range set {
		arr2 = append(arr2, num)
	}

	sort.Ints(arr2)

	dp := map[int]int{-1: 0}

	n, m := len(arr1), len(arr2)
	for _, num := range arr1 {
		newDp := make(map[int]int)

		for prev, ops := range dp {
			// keep
			if num > prev {
				if val, exists := newDp[num]; !exists || ops < val {
					newDp[num] = ops
				}
			}

			// replace
			idx := sort.Search(m, func(i int) bool {
				return arr2[i] > prev
			})

			if idx < m {
				if val, exists := newDp[arr2[idx]]; !exists || ops+1 < val {
					newDp[arr2[idx]] = ops + 1
				}
			}
		}

		dp = newDp
		if len(dp) == 0 { // No valid state found
			return -1
		}
	}

	result := n + 1
	for _, ops := range dp {
		result = min(result, ops)
	}

	if result <= n {
		return result
	}

	return -1
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(n * m * log(m)) where n = len(arr1), m = len(arr2)
// // Space: O(n * m)
// func makeArrayIncreasing(arr1 []int, arr2 []int) int {
// 	set := make(map[int]struct{})
// 	for _, num := range arr2 {
// 		set[num] = struct{}{}
// 	}

// 	arr2 = make([]int, 0, len(set))
// 	for num := range set {
// 		arr2 = append(arr2, num)
// 	}

// 	sort.Ints(arr2)

// 	n, m := len(arr1), len(arr2)
// 	memo := make(map[[2]int]int)

// 	var dfs func(i, prev int) int

// 	dfs = func(i, prev int) int {
// 		if i == n {
// 			return 0
// 		}

// 		state := [2]int{i, prev}
// 		if val, exists := memo[state]; exists {
// 			return val
// 		}

// 		cost := n + 1

// 		// keep
// 		if arr1[i] > prev {
// 			cost = min(cost, dfs(i+1, arr1[i]))
// 		}

// 		// replace
// 		j := sort.Search(len(arr2), func(i int) bool {
// 			return arr2[i] > prev
// 		})

// 		if j < m {
// 			cost = min(cost, 1+dfs(i+1, arr2[j]))
// 		}

// 		memo[state] = cost

// 		return cost
// 	}

// 	ops := dfs(0, -1)
// 	if ops <= n {
// 		return ops
// 	}

// 	return -1
// }
