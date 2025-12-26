package dp

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/

// Approach: Dynamic Programming with State Machine
// dp[i][j][0] = max profit on day i with at most j transactions, not holding stock
// dp[i][j][1] = max profit on day i with at most j transactions, holding stock
// Time: O(n*k)
// Space: O(k)
func maxProfitIV(k int, prices []int) int {
	maxProfitUnlimited := func(prices []int) int {
		profit := 0

		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}

		return profit
	}

	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	// Optimization: if k >= n/2, it's unlimited transactions
	if k >= n/2 {
		return maxProfitUnlimited(prices)
	}

	// dp[j][0] = not holding, dp[j][1] = holding
	buy := make([]int, k+1)
	sell := make([]int, k+1)

	for j := 0; j <= k; j++ {
		buy[j] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := k; j >= 1; j-- {
			sell[j] = max(sell[j], buy[j]+prices[i])
			buy[j] = max(buy[j], sell[j-1]-prices[i])
		}
	}

	return sell[k]
}
