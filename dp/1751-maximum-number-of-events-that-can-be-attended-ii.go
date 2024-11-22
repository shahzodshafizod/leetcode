package dp

import "sort"

// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/

// Approach: Bottom-Up Dynamic Programming + Binary Search
// Time: O(n⋅logn + n⋅k)
// Space: O(n⋅k)
func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	var n = len(events)
	var next = make([]int, n)
	var left, right, mid int
	for idx := 0; idx < n; idx++ {
		left, right = idx+1, n
		for left < right {
			mid = left + (right-left)/2
			if events[mid][0] > events[idx][1] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		next[idx] = right
	}
	var dp = make([][]int, n+1)
	for idx := range dp {
		dp[idx] = make([]int, k+1)
	}
	for idx := n - 1; idx >= 0; idx-- {
		for count := k - 1; count >= 0; count-- {
			dp[idx][count] = max(
				// skip this event
				dp[idx+1][count],
				// attend this event
				events[idx][2]+dp[next[idx]][count+1],
			)
		}
	}
	return dp[0][0]
}

// // Approach: Top-Down Dynamic Programming + Binary Search
// // Time: O(n⋅logn + n⋅k)
// // Space: O(n⋅k)
// func maxValue(events [][]int, k int) int {
// 	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })
// 	var n = len(events)
// 	var next = make([]int, n)
// 	var left, right, mid int
// 	for idx := 0; idx < n; idx++ {
// 		left, right = idx+1, n
// 		for left < right {
// 			mid = left + (right-left)/2
// 			if events[mid][0] > events[idx][1] {
// 				right = mid
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 		next[idx] = right
// 	}
// 	var memo = make([][]*int, n)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, k)
// 	}
// 	var dp func(idx int, count int) int
// 	dp = func(idx int, count int) int {
// 		if idx == n || count == k {
// 			return 0
// 		}
// 		if memo[idx][count] != nil {
// 			return *memo[idx][count]
// 		}
// 		memo[idx][count] = new(int)
// 		*memo[idx][count] = max(
// 			// skip this event
// 			dp(idx+1, count),
// 			// attent this event
// 			events[idx][2]+dp(next[idx], count+1),
// 		)
// 		return *memo[idx][count]
// 	}
// 	return dp(0, 0)
// }
