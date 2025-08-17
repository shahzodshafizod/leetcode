package dp

// In short, you're tasked with finding the likelihood of her final score being between k and n.

// https://leetcode.com/problems/new-21-game/

// Approach #3: Bottom-Up Dynamic Programming (Optimized)
// Time: O(n)
// Space: O(n)
func new21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, n+1)
	dp[0] = 1

	var s float64
	if k > 0 {
		s = 1
	}

	for i := 1; i <= n; i++ {
		dp[i] = s / float64(maxPts)
		if i < k {
			s += dp[i]
		}

		if i-maxPts >= 0 && i-maxPts < k {
			s -= dp[i-maxPts]
		}
	}

	var score float64
	for i := k; i <= n; i++ {
		score += dp[i]
	}

	return score
}

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(n * maxPts)
// // Space: O(n)
// func new21Game(n int, k int, maxPts int) float64 {
// 	dp := make([]float64, n+1)
// 	dp[0] = 1

// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= maxPts; j++ {
// 			if i-j >= 0 && i-j < k {
// 				dp[i] += dp[i-j] / float64(maxPts)
// 			}
// 		}
// 	}

// 	var score float64
// 	for i := k; i <= n; i++ {
// 		score += dp[i]
// 	}

// 	return score
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(k * maxPts)
// // Space: O(k)
// func new21Game(n int, k int, maxPts int) float64 {
// 	memo := make([]*float64, k)

// 	var dp func(int) float64

// 	dp = func(points int) float64 {
// 		if points >= k {
// 			if points <= n {
// 				return 1
// 			}

// 			return 0
// 		}

// 		if memo[points] != nil {
// 			return *memo[points]
// 		}

// 		var score float64
// 		for point := 1; point <= maxPts; point++ {
// 			score += dp(points + point)
// 		}

// 		score /= float64(maxPts)
// 		memo[points] = &score

// 		return score
// 	}

// 	return dp(0)
// }
