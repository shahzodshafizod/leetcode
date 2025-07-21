package fenwicks

import (
	"slices"
)

// https://leetcode.com/problems/count-number-of-teams/

// Approach #4: Binary Indexed Tree (Fenwick Tree)
// Time: O(n⋅log(maxRating))
// Space: O(maxRating)
func numTeams(rating []int) int {
	update := func(bit []int, idx int, delta int) {
		for n := len(bit); idx < n; idx += idx & -idx {
			bit[idx] += delta
		}
	}
	count := func(bit []int, idx int) int {
		count := 0
		for ; idx > 0; idx -= idx & -idx {
			count += bit[idx]
		}
		return count
	}
	maxRating := slices.Max(rating)
	leftBIT := make([]int, maxRating+1)
	rightBIT := make([]int, maxRating+1)
	for _, rating := range rating {
		update(rightBIT, rating, 1)
	}
	teams := 0
	for _, rating := range rating {
		update(rightBIT, rating, -1)
		lSmallers := count(leftBIT, rating-1)
		rSmallers := count(rightBIT, rating-1)
		lGreaters := count(leftBIT, maxRating) - count(leftBIT, rating)
		rGreaters := count(rightBIT, maxRating) - count(rightBIT, rating)
		teams += lSmallers * rGreaters
		teams += lGreaters * rSmallers
		update(leftBIT, rating, 1)
	}
	return teams
}

// // Approach #3: Dynamic Programming (Optimized)
// // Time: (N^2)
// // Space: O(1)
// func numTeams(rating []int) int {
// 	var buildTeams = func(mid int, n int) int {
// 		var minL = 0
// 		for idx := 0; idx < mid; idx++ {
// 			if rating[idx] < rating[mid] {
// 				minL++
// 			}
// 		}
// 		var maxR = 0
// 		for idx := mid + 1; idx < n; idx++ {
// 			if rating[idx] > rating[mid] {
// 				maxR++
// 			}
// 		}
// 		var maxL = mid - minL
// 		var minR = n - mid - 1 - maxR
// 		return minL*maxR + maxL*minR
// 	}
// 	var n = len(rating)
// 	var teams = 0
// 	for mid := 1; mid < n-1; mid++ {
// 		teams += buildTeams(mid, n)
// 	}
// 	return teams
// }

// // Approach #2: Dynamic Programming (Memoization)
// // Time: O(N^2)
// // Space: O(N)
// func numTeams(rating []int) int {
// 	var n = len(rating)
// 	var cache = make([][2][3]int, n)
// 	var backtrack func(idx int, accending int, count int) int
// 	backtrack = func(idx int, accending int, count int) int {
// 		if idx == n {
// 			return 0
// 		}
// 		if count == 3 {
// 			return 1
// 		}
// 		if cache[idx][accending][count] != 0 {
// 			return cache[idx][accending][count]
// 		}
// 		var teams = 0
// 		for j := idx + 1; j < n; j++ {
// 			if accending == 1 && rating[idx] < rating[j] ||
// 				accending == 0 && rating[idx] > rating[j] {
// 				teams += backtrack(j, accending, count+1)
// 			}
// 		}
// 		cache[idx][accending][count] = teams
// 		return teams
// 	}
// 	var teams = 0
// 	for idx := 0; idx < n; idx++ {
// 		teams += backtrack(idx, 1, 1)
// 		teams += backtrack(idx, 0, 1)
// 	}
// 	return teams
// }

// // Approach #1: Brute Force
// // Time: O(N^3)
// // Space: O(1)
// func numTeams(rating []int) int {
// 	var teams = 0
// 	for i, n := 0, len(rating); i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			for k := j + 1; k < n; k++ {
// 				if rating[i] < rating[j] && rating[j] < rating[k] ||
// 					rating[i] > rating[j] && rating[j] > rating[k] {
// 					teams++
// 				}
// 			}
// 		}
// 	}
// 	return teams
// }
