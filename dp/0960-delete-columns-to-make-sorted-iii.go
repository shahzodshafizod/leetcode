package dp

// https://leetcode.com/problems/delete-columns-to-make-sorted-iii/

// Approach #2: Optimized - Dynamic Programming (LIS variant)
// Key insight: Find maximum columns we can KEEP (like LIS on columns)
// dp[i] = max columns we can keep ending at column i
// For each column i, check all previous columns j < i
// If column j can come before i (all chars at j <= chars at i), update dp[i]
// N = number of strings, M = length of each string
// Time: O(M^2 * N) - M^2 column pairs, N to check compatibility
// Space: O(M) - DP array
func minDeletionSize(strs []string) int {
	n, m := len(strs), len(strs[0])

	// dp[i] = maximum number of columns we can keep ending at column i
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1 // Each column can at least keep itself
	}

	// For each column i, check which previous columns j can come before it
	for i := 1; i < m; i++ {
		for j := range i {
			// Check if column j can come before column i
			// This means: for all rows, strs[row][j] <= strs[row][i]
			canPrecede := true

			for row := range n {
				if strs[row][j] > strs[row][i] {
					canPrecede = false

					break
				}
			}

			if canPrecede {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// Maximum columns we can keep
	maxKept := 0
	for _, val := range dp {
		maxKept = max(maxKept, val)
	}

	// Minimum columns to delete
	return m - maxKept
}

// // Approach #1: Brute Force - Try all column subsets
// // Generate all 2^m subsets of columns and check if each subset is valid
// // A subset is valid if keeping only those columns makes each row sorted
// // N = number of strings, M = length of each string
// // Time: O(2^M * N * M) - 2^M subsets, O(N*M) to validate each
// // Space: O(M) - recursion stack and subset tracking
// func minDeletionSize(strs []string) int {
// 	m := len(strs[0])
// 	minDeletions := m // Worst case: delete all columns

// 	isValid := func(columns []int) bool {
// 		// Check if keeping these columns makes all rows sorted
// 		if len(columns) == 0 {
// 			return true
// 		}
// 		// Check each row is sorted with these columns
// 		for _, row := range strs {
// 			for i := 1; i < len(columns); i++ {
// 				if row[columns[i-1]] > row[columns[i]] {
// 					return false
// 				}
// 			}
// 		}
// 		return true
// 	}

// 	var backtrack func(idx int, kept []int)
// 	backtrack = func(idx int, kept []int) {
// 		// Base case: checked all columns
// 		if idx == m {
// 			if isValid(kept) {
// 				deletions := m - len(kept)
// 				minDeletions = min(minDeletions, deletions)
// 			}
// 			return
// 		}

// 		// Try keeping column idx
// 		backtrack(idx+1, append(kept, idx))

// 		// Try deleting column idx
// 		backtrack(idx+1, kept)
// 	}

// 	backtrack(0, []int{})
// 	return minDeletions
// }
