package hashes

// https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/

// Approach: Contribution of Each Character
// Key insight: Instead of checking all substrings, calculate how many substrings
// contain each character occurrence as unique (appears exactly once).
//
// Strategy:
// For each occurrence of character at position i:
// 1. Find previous occurrence of same character (prev)
// 2. Find next occurrence of same character (next)
// 3. Character at i is unique in substrings that:
//   - Start after prev and at or before i
//   - End at or after i and before next
//
// 4. Count = (i - prev) * (next - i)
//
// Example: "ABA"
// - A at index 0: prev=-1, next=2, count=(0-(-1))*(2-0)=2 substrings: "A", "AB"
// - B at index 1: prev=-1, next=3, count=(1-(-1))*(3-1)=4 substrings: "B", "AB", "BA", "ABA"
// - A at index 2: prev=0, next=3, count=(2-0)*(3-2)=2 substrings: "BA", "A"
// Total = 2 + 4 + 2 = 8
//
// Time Complexity: O(n) where n is length of string
// Space Complexity: O(26) = O(1) for storing last occurrence of each character
func uniqueLetterString(s string) int {
	n := len(s)
	result := 0

	// Store indices of each character (A-Z)
	// last[c] = last occurrence index of character c
	last := make([]int, 26)
	for i := range last {
		last[i] = -1
	}

	// Store second-to-last occurrence for each character
	secondLast := make([]int, 26)
	for i := range secondLast {
		secondLast[i] = -1
	}

	for i := range n {
		c := int(s[i] - 'A')

		// Previous occurrence of character c
		prev := secondLast[c]
		// Current occurrence
		curr := last[c]
		// Next occurrence (we don't know yet, so use n)
		next := i

		// If there was a previous occurrence at curr,
		// calculate its contribution now that we know the next occurrence is at i
		if curr != -1 {
			result += (curr - prev) * (next - curr)
		}

		// Update occurrences
		secondLast[c] = last[c]
		last[c] = i
	}

	// Process remaining characters (no next occurrence)
	for c := range 26 {
		prev := secondLast[c]
		curr := last[c]
		next := n

		if curr != -1 {
			result += (curr - prev) * (next - curr)
		}
	}

	return result
}
