package dp

// https://leetcode.com/problems/stone-game-v/

// Approach: Bottom-Up Dynamic Programming
// Time: O(nnn)
// Space: O(nn)
func stoneGameV(stoneValue []int) int {
	var n = len(stoneValue)
	var presum = make([]int, n+1)
	var dp = make([][]int, n)
	for idx := range stoneValue {
		presum[idx+1] = presum[idx] + stoneValue[idx]
		dp[idx] = make([]int, n)
	}
	var right, lsum, rsum int
	for length := 2; length <= n; length++ {
		for left := 0; left <= n-length; left++ {
			right = left + length - 1
			lsum = 0
			for idx := left; idx < right; idx++ {
				lsum += stoneValue[idx]
				rsum = presum[right+1] - presum[idx+1]
				if lsum <= rsum {
					dp[left][right] = max(dp[left][right], lsum+dp[left][idx])
				}
				if rsum <= lsum {
					dp[left][right] = max(dp[left][right], rsum+dp[idx+1][right])
				}
			}
		}
	}
	return dp[0][n-1]
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(nnn)
// // Space: O(nn)
// func stoneGameV(stoneValue []int) int {
// 	var n = len(stoneValue)
// 	var presum = make([]int, n+1)
// 	var cache = make([][]*int, n)
// 	for idx := range stoneValue {
// 		presum[idx+1] = stoneValue[idx] + presum[idx]
// 		cache[idx] = make([]*int, n)
// 	}
// 	var dp func(left int, right int) int
// 	dp = func(left int, right int) int {
// 		if left > right {
// 			return 0
// 		}
// 		if cache[left][right] != nil {
// 			return *cache[left][right]
// 		}
// 		var rsum int
// 		var result, lsum = 0, 0
// 		for idx := left; idx < right; idx++ {
// 			lsum += stoneValue[idx]
// 			rsum = presum[right+1] - presum[idx+1]
// 			if lsum <= rsum {
// 				result = max(result, lsum+dp(left, idx))
// 			}
// 			if rsum <= lsum {
// 				result = max(result, rsum+dp(idx+1, right))
// 			}
// 		}
// 		cache[left][right] = &result
// 		return result
// 	}
// 	return dp(0, n-1)
// }
