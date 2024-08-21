package dp

// https://leetcode.com/problems/strange-printer/

// Bottom Up Dynamic Programming (Tabulation)
// Time: O(N^3)
// Space: O(N^2)
func strangePrinter(s string) int {
	// remove consecutive duplicate characters
	var uniqS = ""
	var prev = ' '
	for _, c := range s {
		if c != prev {
			uniqS += string(c)
		}
		prev = c
	}
	s = uniqS
	var n = len(s)

	var indexes = map[byte]int{}
	var next = make([]int, n)
	for idx := n - 1; idx >= 0; idx-- {
		if indexes[s[idx]] > 0 {
			next[idx] = indexes[s[idx]]
		} else {
			next[idx] = n
		}
		indexes[s[idx]] = idx
	}

	var dp = make([][]int, n+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
		dp[idx][idx] = 1
	}

	for start := n - 2; start >= 0; start-- {
		for end := start + 1; end < n; end++ {
			dp[start][end] = 1 + dp[start+1][end]
			for idx := next[start]; idx <= end; idx = next[idx] {
				dp[start][end] = min(
					dp[start][end],
					dp[start][idx-1]+dp[idx+1][end],
				)
			}
		}
	}

	return dp[0][n-1]
}

// // Top Down Dynamic Programming (Memoization)
// // Time: O(N^3)
// // Space: O(N^2)
// func strangePrinter(s string) int {
// 	uniqS := ""
// 	var prev rune = ' '
// 	for _, c := range s {
// 		if c != prev {
// 			uniqS += string(c)
// 		}
// 		prev = c
// 	}
// 	s = uniqS
// 	var n = len(s)
// 	var memo = make([][]int, n)
// 	for idx := range memo {
// 		memo[idx] = make([]int, n)
// 		memo[idx][idx] = 1
// 	}
// 	var minTurns func(start int, end int) int
// 	minTurns = func(start int, end int) int {
// 		if start > end {
// 			return 0
// 		}
// 		if memo[start][end] != 0 {
// 			return memo[start][end]
// 		}
// 		turns := 1 + minTurns(start+1, end)
// 		for idx := start + 1; idx <= end; idx++ {
// 			if s[idx] == s[start] {
// 				turns = min(turns, minTurns(start, idx-1)+minTurns(idx+1, end))
// 			}
// 		}
// 		memo[start][end] = turns
// 		return turns
// 	}
// 	return minTurns(0, n-1)
// }
