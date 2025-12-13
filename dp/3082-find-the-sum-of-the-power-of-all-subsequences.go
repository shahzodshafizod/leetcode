package dp

// https://leetcode.com/problems/find-the-sum-of-the-power-of-all-subsequences/

// Approach #2: Space-Optimized DP
// Key insight: When we select elements summing to k, remaining elements can be in/out
// dp[v] = sum of 2^(remaining elements) for all ways to achieve sum v
// For each element: dp[v] = 2*dp[v] + dp[v-a] (double existing + add new ways)
// Time: O(n * k), Space: O(k)
func sumOfPower(nums []int, k int) int {
	mod := int(1e9 + 7)
	dp := make([]int, k+1)
	dp[0] = 1

	for _, num := range nums {
		// Iterate backwards to avoid using updated values
		for v := k; v >= 0; v-- {
			// 2 * dp[v]: existing ways, element not included (can be in/out of outer)
			// dp[v - num]: new ways by including this element
			add := 0
			if v >= num {
				add = dp[v-num]
			}

			dp[v] = (2*dp[v] + add) % mod
		}
	}

	return dp[k]
}

// // Approach #1: DFS with Memoization (Brute Force)
// // Try all subsequences, count those with sum = k, power = 2^(n-i)
// // Time: O(n * k) with memoization
// // Space: O(n * k)
// func sumOfPower(nums []int, k int) int {
// 	modPow := func(base, exp, mod int) int {
// 		result := 1
// 		base %= mod

// 		for exp > 0 {
// 			if exp%2 == 1 {
// 				result = (result * base) % mod
// 			}

// 			base = (base * base) % mod
// 			exp /= 2
// 		}

// 		return result
// 	}

// 	mod := int(1e9 + 7)
// 	n := len(nums)
// 	memo := make(map[[2]int]int)

// 	var dfs func(i, sum int) int

// 	dfs = func(i, sum int) int {
// 		if sum == k {
// 			// Remaining n-i elements can be in or out of outer subsequence
// 			return modPow(2, n-i, mod)
// 		}

// 		if sum > k || i >= n {
// 			return 0
// 		}

// 		key := [2]int{i, sum}
// 		if val, ok := memo[key]; ok {
// 			return val
// 		}
// 		// Skip: element not included (can be in/out of outer) = 2 * dfs(i+1, sum)
// 		// Take: element included in inner subsequence = dfs(i+1, sum+nums[i])
// 		result := (2*dfs(i+1, sum) + dfs(i+1, sum+nums[i])) % mod
// 		memo[key] = result

// 		return result
// 	}

// 	return dfs(0, 0)
// }
