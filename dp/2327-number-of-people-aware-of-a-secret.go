package dp

// https://leetcode.com/problems/number-of-people-aware-of-a-secret/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(n)
// Space: O(n)
func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = int(1e9 + 7)

	dp := make([]int, n+1)
	dp[1] = 1

	var people int

	for day := 2; day <= n; day++ {
		if day-delay > 0 {
			people = (people + dp[day-delay]) % mod
		}

		if day-forget > 0 {
			people = (people - dp[day-forget] + mod) % mod
		}

		dp[day] = people
	}

	var total int
	for day := n - forget + 1; day <= n; day++ {
		total = (total + dp[day]) % mod
	}

	return total
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(nn)
// // Space: O(n)
// func peopleAwareOfSecret(n int, delay int, forget int) int {
// 	const mod = int(1e9 + 7)

// 	memo := make([]*int, n+1)

// 	var dp func(day int) int

// 	dp = func(day int) int {
// 		if memo[day] != nil {
// 			return *memo[day]
// 		}

// 		var people int
// 		if day+forget-1 >= n {
// 			// if this person contributes to the last day
// 			people = 1
// 		}

// 		for nxt := delay; nxt < forget; nxt++ {
// 			if day+nxt <= n {
// 				people = (people + dp(day+nxt)) % mod
// 			}
// 		}

// 		memo[day] = &people

// 		return people
// 	}

// 	return dp(1)
// }
