package graphs

import "math"

// https://leetcode.com/problems/cheapest-flights-within-k-stops/

// Bellman Ford's Algorithm
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = math.MaxInt
	}

	prices[src] = 0
	for i := 0; i <= k; i++ {
		tmpPrices := append([]int{}, prices...)

		for _, flight := range flights {
			from := flight[0]
			if prices[from] == math.MaxInt {
				continue
			}

			to := flight[1]
			price := flight[2]

			if prices[from]+price < tmpPrices[to] {
				tmpPrices[to] = prices[from] + price
			}
		}

		prices = tmpPrices
	}

	if prices[dst] == math.MaxInt {
		prices[dst] = -1
	}

	return prices[dst]
}
