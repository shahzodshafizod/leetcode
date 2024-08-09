package greedy

import "sort"

// https://leetcode.com/problems/most-profit-assigning-work/

// // Approach #5: Memoization
// // time: O(n + m + maxAbility)
// // space: O(maxAbility)
// func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
// 	var maxAbility = slices.Max(worker)
// 	var jobs = make([]int, maxAbility+1)
// 	for idx, n := 0, len(profit); idx < n; idx++ { // O(n)
// 		if difficulty[idx] <= maxAbility {
// 			jobs[difficulty[idx]] = max(jobs[difficulty[idx]], profit[idx])
// 		}
// 	}
// 	for ability := 1; ability <= maxAbility; ability++ { // O(maxAbility)
// 		jobs[ability] = max(jobs[ability], jobs[ability-1])
// 	}
// 	var netProfit = 0
// 	for _, ability := range worker { // O(m)
// 		netProfit += jobs[ability]
// 	}
// 	return netProfit
// }

// Approach #4: Greedy and Two-Pointers
// time: O(nlogn + mlogm)
// space: O(2n) = O(n)
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var n = len(difficulty)
	var jobProfile = make([][2]int, n)
	for idx := 0; idx < n; idx++ {
		jobProfile[idx][0] = difficulty[idx]
		jobProfile[idx][1] = profit[idx]
	}
	sort.Slice(jobProfile, func(i, j int) bool { // O(n log n)
		return jobProfile[i][0] < jobProfile[j][0]
	})
	sort.Ints(worker) // O(m log m)
	var netProfit = 0
	var idx, maxProfit = 0, 0
	for _, ability := range worker { // O(max(m,n))
		for idx < n && jobProfile[idx][0] <= ability {
			maxProfit = max(maxProfit, jobProfile[idx][1])
			idx++
		}
		netProfit += maxProfit
	}
	return netProfit
}

// // Approach #3: Binary Search and Greedy (Sort by profit)
// // time: O(mlogn + nlogn)
// // space: O(2n) = O(n)
// func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
// 	var n = len(difficulty)
// 	var jobProfile = make([][2]int, n) // [profit,difficulty]
// 	for idx := range difficulty {      // O(n)
// 		jobProfile[idx][0] = profit[idx]
// 		jobProfile[idx][1] = difficulty[idx]
// 	}
// 	sort.Slice(jobProfile, func(i, j int) bool { // O(n log n)
// 		return jobProfile[i][0] > jobProfile[j][0]
// 	})
// 	for idx := 1; idx < n; idx++ { // O(n)
// 		jobProfile[idx][1] = min(jobProfile[idx][1], jobProfile[idx-1][1])
// 	}
// 	var netProfit = 0
// 	var maxProfit, left, right, mid int
// 	for _, ability := range worker { // O(m)
// 		left, right = 0, n-1
// 		maxProfit = 0
// 		for left <= right { // O(log n)
// 			mid = (left + right) / 2
// 			if jobProfile[mid][1] <= ability {
// 				maxProfit = max(maxProfit, jobProfile[mid][0])
// 				right = mid - 1
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 		netProfit += maxProfit
// 	}
// 	return netProfit
// }

// // Approach #2: Binary Search and Greedy (Sort by Job Difficulty)
// // time: O(mlogn + nlogn)
// // space: O(2n) = O(n)
// func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
// 	var n = len(difficulty)
// 	var jobProfile = make([][2]int, n) // [difficulty,profit]
// 	for idx := range difficulty {      // O(n)
// 		jobProfile[idx][0] = difficulty[idx]
// 		jobProfile[idx][1] = profit[idx]
// 	}
// 	sort.Slice(jobProfile, func(i, j int) bool { // O(n log n)
// 		return jobProfile[i][0] < jobProfile[j][0]
// 	})
// 	for idx := 1; idx < n; idx++ { // O(n)
// 		jobProfile[idx][1] = max(jobProfile[idx][1], jobProfile[idx-1][1])
// 	}
// 	var netProfit = 0
// 	var maxProfit, left, right, mid int
// 	for _, ability := range worker { // O(m)
// 		left, right = 0, n-1
// 		maxProfit = 0
// 		for left <= right { // O(log n)
// 			mid = (left + right) / 2
// 			if jobProfile[mid][0] <= ability {
// 				maxProfit = jobProfile[mid][1]
// 				left = mid + 1
// 			} else {
// 				right = mid - 1
// 			}
// 		}
// 		netProfit += maxProfit
// 	}
// 	return netProfit
// }

// // Approach #1: Brute Force
// // time: O(mxn)
// // space: O(1)
// func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
// 	var netprofit = 0
// 	var maxProfit int
// 	for _, ability := range worker {
// 		maxProfit = 0
// 		for idx := range difficulty {
// 			if difficulty[idx] <= ability && profit[idx] > maxProfit {
// 				maxProfit = profit[idx]
// 			}
// 		}
// 		netprofit += maxProfit
// 	}
// 	return netprofit
// }
