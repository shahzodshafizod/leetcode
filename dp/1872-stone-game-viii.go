package dp

// https://leetcode.com/problems/stone-game-viii/

// Approach #4: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(n)
// Space: O(n)
func stoneGameVIII(stones []int) int {
	n := len(stones)
	presum := make([]int, n)
	presum[0] = stones[0]

	for idx := 1; idx < n; idx++ {
		presum[idx] = presum[idx-1] + stones[idx]
	}

	dp := presum[n-1]
	for idx := n - 2; idx > 0; idx-- {
		dp = max(dp, presum[idx]-dp)
	}

	return dp
}

// // Approach #3: Bottom-Up Dynamic Programming
// // Time: O(n)
// // Space: O(n)
// func stoneGameVIII(stones []int) int {
// 	var n = len(stones)
// 	var presum = make([]int, n)
// 	presum[0] = stones[0]
// 	for idx := 1; idx < n; idx++ {
// 		presum[idx] = presum[idx-1] + stones[idx]
// 	}
// 	var dp = make([]int, n)
// 	dp[n-1] = presum[n-1]
// 	for idx := n - 2; idx > 0; idx-- {
// 		dp[idx] = max(dp[idx+1], presum[idx]-dp[idx+1])
// 	}
// 	return dp[1]
// }

// // Approach #2: Top-Down Dynamic Programming (Time-Optimized)
// // Time: O(n)
// // Space: O(n)
// func stoneGameVIII(stones []int) int {
// 	var n = len(stones)
// 	var presum = make([]int, n)
// 	presum[0] = stones[0]
// 	for idx := 1; idx < n; idx++ {
// 		presum[idx] = presum[idx-1] + stones[idx]
// 	}
// 	var memo = make([]*int, n)
// 	var dp func(idx int) int
// 	dp = func(idx int) int {
// 		if idx == n-1 {
// 			return presum[idx]
// 		}
// 		if memo[idx] != nil {
// 			return *memo[idx]
// 		}
// 		memo[idx] = new(int)
// 		*memo[idx] = max(dp(idx+1), presum[idx]-dp(idx+1))
// 		return *memo[idx]
// 	}
// 	return dp(1)
// }

// // Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
// // Time: O(nn)
// // Space: O(n)
// func stoneGameVIII(stones []int) int {
// 	var n = len(stones)
// 	var presum = make([]int, n)
// 	presum[0] = stones[0]
// 	for idx := 1; idx < n; idx++ {
// 		presum[idx] = presum[idx-1] + stones[idx]
// 	}
// 	var memo = make([]*int, n)
// 	var dp func(idx int) int
// 	dp = func(idx int) int {
// 		if idx == n {
// 			return 0
// 		}
// 		if memo[idx] != nil {
// 			return *memo[idx]
// 		}
// 		var diff = math.MinInt
// 		for x := idx; x < n; x++ {
// 			diff = max(diff, presum[x]-dp(x+1))
// 		}
// 		memo[idx] = &diff
// 		return diff
// 	}
// 	return dp(1)
// }
