package dp

// https://leetcode.com/problems/coin-change/

// Bottom-Up
// time: O(amount*len(coins))
// space: O(amount)
func coinChange(coins []int, amount int) int {
	counts := make([]int, amount+1)
	counts[0] = 0

	for a := 1; a <= amount; a++ {
		counts[a] = amount + 1
		for _, coin := range coins {
			if a-coin >= 0 && counts[a-coin] != amount+1 {
				counts[a] = min(counts[a], 1+counts[a-coin])
			}
		}
	}

	if counts[amount] == amount+1 {
		counts[amount] = -1
	}

	return counts[amount]
}

// // Time Limit Exceeded
// // Top-Down
// func coinChange(coins []int, amount int) int {
// 	var minCount = math.MaxInt
// 	var dfs func(idx int, amount int, count int)
// 	dfs = func(idx int, amount int, count int) {
// 		if idx < 0 || amount < 0 || count > minCount {
// 			return
// 		}
// 		if amount == 0 {
// 			minCount = min(minCount, count)
// 			return
// 		}
// 		dfs(idx, amount-coins[idx], count+1)
// 		dfs(idx-1, amount, count)
// 	}
// 	dfs(len(coins)-1, amount, 0)
// 	if minCount == math.MaxInt {
// 		minCount = -1
// 	}
// 	return minCount
// }
