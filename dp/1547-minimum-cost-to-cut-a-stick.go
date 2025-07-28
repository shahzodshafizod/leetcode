package dp

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-cost-to-cut-a-stick/

// Approach: Bottom-Up Dynamic Programming
// Time: O(mmm), m=len(cuts)
// Space: O(mm)
func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)
	m := len(cuts)
	dp := make([][]int, m)

	for idx := range dp {
		dp[idx] = make([]int, m)
	}

	var start, end, mid, cost int
	for diff := 2; diff < m; diff++ {
		for start = 0; start < m-diff; start++ {
			end = start + diff
			dp[start][end] = math.MaxInt
			cost = cuts[end] - cuts[start]

			for mid = start + 1; mid < end; mid++ {
				dp[start][end] = min(dp[start][end],
					cost+dp[start][mid]+dp[mid][end],
				)
			}
		}
	}

	return dp[0][m-1]
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(mmm), m=len(cuts)
// // Space: O(mm)
// func minCost(n int, cuts []int) int {
// 	cuts = append(cuts, 0, n)
// 	sort.Ints(cuts)
// 	var m = len(cuts)
// 	var memo = make([][]*int, m)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, m)
// 	}
// 	var dp func(start int, end int) int
// 	dp = func(start int, end int) int {
// 		if end-start == 1 {
// 			return 0
// 		}
// 		if memo[start][end] != nil {
// 			return *memo[start][end]
// 		}
// 		var best, cost = math.MaxInt, cuts[end] - cuts[start]
// 		for mid := start + 1; mid < end; mid++ {
// 			best = min(best, cost+dp(start, mid)+dp(mid, end))
// 		}
// 		memo[start][end] = &best
// 		return best
// 	}
// 	return dp(0, m-1)
// }
