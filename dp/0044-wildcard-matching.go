package dp

// https://leetcode.com/problems/wildcard-matching/

// Approach: Optimized - DP with tabulation
// Use 2D DP table where dp[i][j] = whether s[0:i] matches p[0:j]
// Handle '?' and '*' with specific transitions
// N = length of s, M = length of p
// Time: O(N * M) - fill DP table
// Space: O(N * M) - DP table
func isMatchWildcard(s string, p string) bool {
	n, m := len(s), len(p)

	// dp[i][j] = true if s[0:i] matches p[0:j]
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	// Empty string matches empty pattern
	dp[0][0] = true

	// Empty string can match pattern with only '*'
	for j := 1; j <= m; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Fill DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			switch p[j-1] {
			case '*':
				// '*' matches empty sequence or one+ characters
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			case '?', s[i-1]:
				// '?' or exact character match
				dp[i][j] = dp[i-1][j-1]
			default:
			}
			// else: no match, remains false
		}
	}

	return dp[n][m]
}
