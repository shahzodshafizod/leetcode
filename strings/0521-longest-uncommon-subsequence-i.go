package strings

// https://leetcode.com/problems/longest-uncommon-subsequence-i/

// Approach 1: Brute Force - Generate All Subsequences
// Generate all subsequences of both strings
// Find longest subsequence that appears in only one string
// Time: O(2^n * 2^m) - exponential, not practical
// Space: O(2^n + 2^m)

// Approach 2: Brain Teaser / Logic (Optimal)
// Key insight: If strings are different, the longer string (or either if same length)
// is an uncommon subsequence. If strings are identical, no uncommon subsequence exists.
//
// Why? If a != b, then "a" itself is a subsequence of "a" but not of "b".
// If a == b, then every subsequence of a is also a subsequence of b.
//
// Time: O(min(len(a), len(b))) - for string comparison
// Space: O(1)

func findLUSlength(a string, b string) int {
	// If strings are identical, no uncommon subsequence exists
	if a == b {
		return -1
	}

	// If strings are different, the longer one (or either) is uncommon
	// Return the maximum length
	if len(a) > len(b) {
		return len(a)
	}

	return len(b)
}
