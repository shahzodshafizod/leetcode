package dp

// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/

// Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(S⋅min(S,A)), S=steps, A=arrLen
// Space: O(min(S,A))
func numWays(steps int, arrLen int) int {
	arrLen = min(arrLen, steps) // cannot move further than "steps" steps
	var prev, curr = make([]int, arrLen), make([]int, arrLen)
	curr[0] = 1
	const MOD int = 1e9 + 7
	for st := 1; st <= steps; st++ {
		prev, curr = curr, prev
		for pos := arrLen - 1; pos >= 0; pos-- {
			curr[pos] = prev[pos]
			if pos > 0 {
				curr[pos] = (curr[pos] + prev[pos-1]) % MOD
			}
			if pos+1 < arrLen {
				curr[pos] = (curr[pos] + prev[pos+1]) % MOD
			}
		}
	}
	return curr[0]
}

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(S⋅min(S,A)), S=steps, A=arrLen
// // Space: O(S⋅min(S,A))
// func numWays(steps int, arrLen int) int {
// 	arrLen = min(arrLen, steps) // cannot move further than "steps" steps
// 	var dp = make([][]int, arrLen)
// 	for idx := range dp {
// 		dp[idx] = make([]int, steps+1)
// 	}
// 	dp[0][0] = 1
// 	const MOD int = 1e9 + 7
// 	for st := 1; st <= steps; st++ {
// 		for pos := arrLen - 1; pos >= 0; pos-- {
// 			dp[pos][st] = dp[pos][st-1]
// 			if pos > 0 {
// 				dp[pos][st] = (dp[pos][st] + dp[pos-1][st-1]) % MOD
// 			}
// 			if pos+1 < arrLen {
// 				dp[pos][st] = (dp[pos][st] + dp[pos+1][st-1]) % MOD
// 			}
// 		}
// 	}
// 	return dp[0][steps]
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(S⋅min(S,A)), S=steps, A=arrLen
// // Space: O(S⋅min(S,A))
// func numWays(steps int, arrLen int) int {
// 	arrLen = min(arrLen, steps) // cannot move further than "steps" steps
// 	var memo = make([][]*int, arrLen)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, steps+1)
// 	}
// 	var dp func(pos int, steps int) int
// 	dp = func(pos int, steps int) int {
// 		if pos < 0 || pos == arrLen {
// 			return 0
// 		}
// 		if steps == 0 {
// 			if pos == 0 {
// 				return 1
// 			}
// 			return 0
// 		}
// 		if memo[pos][steps] != nil {
// 			return *memo[pos][steps]
// 		}
// 		memo[pos][steps] = new(int)
// 		*memo[pos][steps] = (dp(pos-1, steps-1) +
// 			dp(pos, steps-1) +
// 			dp(pos+1, steps-1)) % int(1e9+7)
// 		return *memo[pos][steps]
// 	}
// 	return dp(0, steps)
// }
