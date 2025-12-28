package dp

// https://leetcode.com/problems/maximum-profit-from-trading-stocks-with-discounts/

// Approach: Tree DP with Bounded Knapsack
// Intuition: Use tree DP where for each node we track:
// - dp0[cost] = max profit in subtree using exact cost when parent didn't buy
// - dp1[cost] = max profit in subtree using exact cost when parent bought (discount available)
// We use bounded arrays sized by maximum possible cost in each subtree.
// Time Complexity: O(n * budget^2)
// Space Complexity: O(n * budget)
func maximumProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	// Build adjacency list
	children := make([][]int, n+1)
	for _, edge := range hierarchy {
		children[edge[0]] = append(children[edge[0]], edge[1])
	}

	type result struct {
		dp0  []int
		dp1  []int
		size int
	}

	var dfs func(int) result

	dfs = func(u int) result {
		cost := present[u-1]
		dCost := cost / 2
		profit := future[u-1] - cost
		dProfit := future[u-1] - dCost

		// subProfit0[w] = max profit from children with exact cost w, children see unbought parent
		// subProfit1[w] = max profit from children with exact cost w, children see bought parent
		subProfit0 := []int{0}
		subProfit1 := []int{0}

		uSize := cost

		for _, v := range children[u] {
			childRes := dfs(v)
			uSize += childRes.size

			// Merge with child using bounded knapsack
			newLen := min(budget+1, len(subProfit0)+len(childRes.dp0)-1)

			newSub0 := make([]int, newLen)
			for i := range newSub0 {
				newSub0[i] = -1
			}

			for i := 0; i < len(subProfit0) && i < newLen; i++ {
				if subProfit0[i] == -1 {
					continue
				}

				for j := 0; j < len(childRes.dp0) && i+j < newLen; j++ {
					if childRes.dp0[j] != -1 {
						newSub0[i+j] = max(newSub0[i+j], subProfit0[i]+childRes.dp0[j])
					}
				}
			}

			subProfit0 = newSub0

			newLen = min(budget+1, len(subProfit1)+len(childRes.dp1)-1)

			newSub1 := make([]int, newLen)
			for i := range newSub1 {
				newSub1[i] = -1
			}

			for i := 0; i < len(subProfit1) && i < newLen; i++ {
				if subProfit1[i] == -1 {
					continue
				}

				for j := 0; j < len(childRes.dp1) && i+j < newLen; j++ {
					if childRes.dp1[j] != -1 {
						newSub1[i+j] = max(newSub1[i+j], subProfit1[i]+childRes.dp1[j])
					}
				}
			}

			subProfit1 = newSub1
		}

		// Build dp0: parent didn't buy
		maxLen0 := min(budget+1, max(len(subProfit0), len(subProfit1)+cost))

		dp0 := make([]int, maxLen0)
		for i := range dp0 {
			dp0[i] = -1
		}

		// Option 1: don't buy u
		for i := 0; i < len(subProfit0) && i < maxLen0; i++ {
			dp0[i] = max(dp0[i], subProfit0[i])
		}

		// Option 2: buy u at full cost
		if cost < maxLen0 {
			for i := 0; i < len(subProfit1) && i+cost < maxLen0; i++ {
				if subProfit1[i] != -1 {
					dp0[i+cost] = max(dp0[i+cost], subProfit1[i]+profit)
				}
			}
		}

		// Build dp1: parent bought (discount available)
		maxLen1 := min(budget+1, max(len(subProfit0), len(subProfit1)+dCost))

		dp1 := make([]int, maxLen1)
		for i := range dp1 {
			dp1[i] = -1
		}

		// Option 1: don't buy u
		for i := 0; i < len(subProfit0) && i < maxLen1; i++ {
			dp1[i] = max(dp1[i], subProfit0[i])
		}

		// Option 2: buy u at discounted cost
		if dCost < maxLen1 {
			for i := 0; i < len(subProfit1) && i+dCost < maxLen1; i++ {
				if subProfit1[i] != -1 {
					dp1[i+dCost] = max(dp1[i+dCost], subProfit1[i]+dProfit)
				}
			}
		}

		return result{dp0, dp1, uSize}
	}

	res := dfs(1)

	ans := 0
	for i := 0; i <= budget && i < len(res.dp0); i++ {
		ans = max(ans, res.dp0[i])
	}

	return ans
}
