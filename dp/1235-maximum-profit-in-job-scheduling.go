package dp

import (
	"sort"
)

// https://leetcode.com/problems/maximum-profit-in-job-scheduling/

// Approach: Bottom-Up Dynamic Programming + Binary Search
// Time: O(n⋅logn)
// Space: O(n)
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var n = len(startTime)
	var jobs = make([][3]int, n)
	for idx := range jobs {
		jobs[idx] = [3]int{startTime[idx], endTime[idx], profit[idx]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	var dp = make([]int, n+1)
	var left, right, mid, next int
	for idx := n - 1; idx >= 0; idx-- {
		left, right = idx+1, n
		for left < right {
			mid = left + (right-left)/2
			if jobs[mid][0] >= jobs[idx][1] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		next = right
		dp[idx] = max(
			dp[idx+1],             // skip this job
			jobs[idx][2]+dp[next], // take this job
		)
	}
	return dp[0]
}

// // Approach: Top-Down Dynamic Programming + Binary Search
// // Time: O(n⋅logn)
// // Space: O(n)
// func jobScheduling(startTime []int, endTime []int, profit []int) int {
// 	var n = len(startTime)
// 	var jobs = make([][3]int, n)
// 	for idx := range jobs {
// 		jobs[idx] = [3]int{startTime[idx], endTime[idx], profit[idx]}
// 	}
// 	sort.Slice(jobs, func(i, j int) bool {
// 		return jobs[i][0] < jobs[j][0]
// 	})
// 	var memo = make([]*int, n)
// 	var dp func(idx int) int
// 	dp = func(idx int) int {
// 		if idx == n {
// 			return 0
// 		}
// 		if memo[idx] != nil {
// 			return *memo[idx]
// 		}
// 		var left, right = idx + 1, n
// 		var mid int
// 		for left < right {
// 			mid = left + (right-left)/2
// 			if jobs[mid][0] >= jobs[idx][1] {
// 				right = mid
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 		var next = right
// 		memo[idx] = new(int)
// 		*memo[idx] = max(
// 			dp(idx+1),             // skip this job
// 			jobs[idx][2]+dp(next), // take this job
// 		)
// 		return *memo[idx]
// 	}
// 	return dp(0)
// }
