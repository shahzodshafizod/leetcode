package dp

// https://leetcode.com/problems/frog-jump/

/*
EARLY EXIT:
if (n * (n - 1) // 2) < stones[n - 1]:
    return False
If the frog makes n jumps and always increases its steps by 1,
the total distance is: 1+2+3+...+n = n(n+1)/2
However, the frog needs n-1 steps to reach the last stone: (n-1)n/2
Means: If the last stone is farther than the maximum possible distance the frog could travel
*/

// Approach #1: Top-Down Dynamic Programming
// Time: O(nn)
// Space: O(nn)
func canCross(stones []int) bool {
	n := len(stones)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][1] = true

	var k int

	for i := 1; i < n; i++ {
		for j := range i {
			k = stones[i] - stones[j]
			if k <= n && dp[j][k] {
				if i == n-1 {
					return true
				}

				dp[i][k-1] = true
				dp[i][k] = true
				dp[i][k+1] = true
			}
		}
	}

	return false
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func canCross(stones []int) bool {
// 	n := len(stones)
// 	if n*(n-1)/2 < stones[n-1] {
// 		return false
// 	}

// 	indices := make(map[int]int, n)

// 	seen := make([][]bool, n)
// 	for i := range n {
// 		indices[stones[i]] = i + 1
// 		seen[i] = make([]bool, n)
// 	}

// 	var dp func(i int, k int) bool

// 	dp = func(i, k int) bool {
// 		if i == n-1 {
// 			return true
// 		}

// 		if seen[i][k] {
// 			return false
// 		}

// 		j := indices[stones[i]+k] - 1
// 		if j <= i {
// 			return false
// 		}

// 		if dp(j, k-1) || dp(j, k) || dp(j, k+1) {
// 			return true
// 		}

// 		seen[i][k] = true

// 		return false
// 	}

// 	return dp(0, 1)
// }
