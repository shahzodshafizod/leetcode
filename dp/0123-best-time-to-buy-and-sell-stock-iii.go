package dp

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/

// // Time: O(N^2)
// // Space: O(1)
// func maxProfitIII(prices []int) int {
// 	var getProfit = func(start int, end int) int {
// 		var minPrice = 100000
// 		var maxProfit = 0
// 		for idx := start; idx <= end; idx++ {
// 			minPrice = min(minPrice, prices[idx])
// 			maxProfit = max(maxProfit, prices[idx]-minPrice)
// 		}
// 		return maxProfit
// 	}
// 	var n = len(prices)
// 	var profit = 0
// 	for idx := 0; idx < n; idx++ {
// 		profit = max(profit, getProfit(0, idx)+getProfit(idx+1, n-1))
// 	}
// 	return profit
// }

// // Time: O(3N) = O(N)
// // Space: O(2N) = O(N)
// func maxProfitIII(prices []int) int {
// 	var n = len(prices)

// 	var leftProfits = make([]int, n)
// 	var minPrice = 100000
// 	var maxProfit = 0
// 	for idx := range prices {
// 		minPrice = min(minPrice, prices[idx])
// 		maxProfit = max(maxProfit, prices[idx]-minPrice)
// 		leftProfits[idx] = maxProfit
// 	}

// 	var rightProfits = make([]int, n)
// 	var maxPrice = 0
// 	maxProfit = 0
// 	for idx := n - 1; idx >= 0; idx-- {
// 		maxPrice = max(maxPrice, prices[idx])
// 		maxProfit = max(maxProfit, maxPrice-prices[idx])
// 		rightProfits[idx] = maxProfit
// 	}

// 	var profit = 0
// 	for idx := 0; idx < n; idx++ {
// 		profit = max(profit, leftProfits[idx]+rightProfits[idx])
// 	}
// 	return profit
// }

// // Time: O(2N) = O(N)
// // Space: O(N)
// func maxProfitIII(prices []int) int {
// 	var n = len(prices)

// 	var leftProfits = make([]int, n)
// 	var minPrice = 100000
// 	var maxProfit = 0
// 	for idx := range prices {
// 		minPrice = min(minPrice, prices[idx])
// 		maxProfit = max(maxProfit, prices[idx]-minPrice)
// 		leftProfits[idx] = maxProfit
// 	}

// 	var maxPrice = 0
// 	maxProfit = 0
// 	var result = 0
// 	for idx := n - 1; idx >= 0; idx-- {
// 		maxPrice = max(maxPrice, prices[idx])
// 		maxProfit = max(maxProfit, maxPrice-prices[idx])
// 		result = max(result, leftProfits[idx]+maxProfit)
// 	}

// 	return result
// }

// Time: O(N)
// Space: O(1)
func maxProfitIII(prices []int) int {
	buyFirst, sellFirst := 100000, 0
	buySecond, sellSecond := 100000, 0
	for _, price := range prices {
		buyFirst = min(buyFirst, price)
		sellFirst = max(sellFirst, price-buyFirst)
		buySecond = min(buySecond, price-sellFirst)
		sellSecond = max(sellSecond, price-buySecond)
	}
	return sellSecond
}

// // Time: O(3N) = O(N)
// // Space: O(1)
// func maxProfitIII(prices []int) int {
// 	var profit = func(prices []int) (int, int, int) {
// 		var maxProfit, minIndex, start, end = 0, 0, 0, -1
// 		for idx := range prices {
// 			if prices[idx] < prices[minIndex] {
// 				minIndex = idx
// 			}
// 			profit := prices[idx] - prices[minIndex]
// 			if maxProfit < profit {
// 				maxProfit = profit
// 				end = idx
// 				start = minIndex
// 			}
// 		}
// 		return maxProfit, start, end
// 	}

// 	var loss = func(prices []int) int {
// 		loss := 0
// 		maxPrice := prices[0]
// 		for _, price := range prices {
// 			maxPrice = max(price, maxPrice)
// 			loss = min(price-maxPrice, loss)
// 		}
// 		return -loss
// 	}

// 	maxProfit, start, end := profit(prices)
// 	if end == -1 {
// 		return 0
// 	}

// 	leftProfit, _, _ := profit(prices[:start])
// 	rightProfit, _, _ := profit(prices[end+1:])
// 	midLoss := loss(prices[start : end+1])

// 	return maxProfit + max(leftProfit, rightProfit, midLoss)
// }
