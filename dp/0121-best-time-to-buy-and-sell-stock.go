package dp

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	profit := 0

	cheapest := 10000
	for _, price := range prices {
		cheapest = min(cheapest, price)
		profit = max(profit, price-cheapest)
	}

	return profit
}
