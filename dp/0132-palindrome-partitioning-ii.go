package dp

// https://leetcode.com/problems/palindrome-partitioning-ii/

// Approach: Optimized - DP with palindrome pre-computation
// Pre-compute palindrome table, then use DP for minimum cuts
// dp[i] = minimum cuts needed for s[0:i+1]
// isPalindrome[i][j] = true if s[i:j+1] is palindrome
// N = length of string
// Time: O(N^2) - palindrome table construction + DP
// Space: O(N^2) - palindrome table + O(N) for DP array
func minCut(s string) int {
	n := len(s)

	// Build palindrome table
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	// Every single character is a palindrome
	for i := range n {
		isPalindrome[i][i] = true
	}

	// Check for length 2 and more
	for length := 2; length <= n; length++ {
		for start := 0; start <= n-length; start++ {
			end := start + length - 1
			if s[start] == s[end] {
				if length == 2 {
					isPalindrome[start][end] = true
				} else {
					isPalindrome[start][end] = isPalindrome[start+1][end-1]
				}
			}
		}
	}

	// DP array: dp[i] = minimum cuts for s[0:i+1]
	dp := make([]int, n)

	for i := range n {
		if isPalindrome[0][i] {
			// If entire substring from start is palindrome, no cuts needed
			dp[i] = 0
		} else {
			// Try all possible cuts
			minCuts := i // worst case: i cuts for i+1 characters
			for j := range i {
				// If s[j+1:i+1] is palindrome, we can cut at position j
				if isPalindrome[j+1][i] {
					if dp[j]+1 < minCuts {
						minCuts = dp[j] + 1
					}
				}
			}

			dp[i] = minCuts
		}
	}

	return dp[n-1]
}
