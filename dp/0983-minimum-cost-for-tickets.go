package dp

import "math"

// https://leetcode.com/problems/minimum-cost-for-tickets/

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	memo := make([]*int, n)

	var dp func(idx int) int

	dp = func(idx int) int {
		if idx == n {
			return 0
		}

		key := idx
		if memo[key] != nil {
			return *memo[key]
		}

		idx++
		day := 1
		cost := math.MaxInt

		for c, pass := range []int{1, 7, 30} {
			for ; idx < n && day+(days[idx]-days[idx-1]) <= pass; idx++ {
				day += days[idx] - days[idx-1]
			}

			cost = min(cost, costs[c]+dp(idx))
		}

		memo[key] = &cost

		return cost
	}

	return dp(0)
}
