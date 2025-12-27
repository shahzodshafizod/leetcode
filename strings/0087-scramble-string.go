package strings

import "slices"

// https://leetcode.com/problems/scramble-string/

// Approach: Dynamic Programming with Memoization (Top-Down)
// Time: O(n^4) - for each substring pair, try all split points
// Space: O(n^3) - memoization cache
func isScramble(s1 string, s2 string) bool {
	// Helper function to check if two strings have the same characters
	sameChars := func(s1, s2 string) bool {
		if len(s1) != len(s2) {
			return false
		}

		// Sort and compare
		r1 := []rune(s1)
		r2 := []rune(s2)

		slices.Sort(r1)
		slices.Sort(r2)

		for i := range r1 {
			if r1[i] != r2[i] {
				return false
			}
		}

		return true
	}

	// Memoization cache: "s1,s2" -> bool
	memo := make(map[string]bool)

	var helper func(s1, s2 string) bool

	helper = func(s1, s2 string) bool {
		// Base cases
		if s1 == s2 {
			return true
		}

		// Quick check: if character frequencies don't match, can't be scramble
		if !sameChars(s1, s2) {
			return false
		}

		// Check memo
		key := s1 + "," + s2
		if val, exists := memo[key]; exists {
			return val
		}

		// Try all possible split points
		n := len(s1)
		for i := 1; i < n; i++ {
			// Case 1: No swap
			// s1[:i] matches s2[:i] AND s1[i:] matches s2[i:]
			if helper(s1[:i], s2[:i]) && helper(s1[i:], s2[i:]) {
				memo[key] = true

				return true
			}

			// Case 2: Swap
			// s1[:i] matches s2[n-i:] AND s1[i:] matches s2[:n-i]
			if helper(s1[:i], s2[n-i:]) && helper(s1[i:], s2[:n-i]) {
				memo[key] = true

				return true
			}
		}

		memo[key] = false

		return false
	}

	return helper(s1, s2)
}
