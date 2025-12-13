package strings

// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/

// Approach #2: Optimized - Use Z-algorithm for efficient prefix matching
// Z-algorithm builds an array where z[i] is the length of longest substring
// starting at i that matches a prefix of the string
// Time: O(n) for Z-algorithm construction
// Space: O(n) for Z array
func minimumTimeToInitialState(word string, k int) int {
	n := len(word)

	// Build Z-array
	zAlgorithm := func(s string) []int {
		n := len(s)
		z := make([]int, n)
		z[0] = n
		left, right := 0, 0

		for i := 1; i < n; i++ {
			if i > right {
				left, right = i, i
				for right < n && s[right-left] == s[right] {
					right++
				}

				z[i] = right - left
				right--
			} else {
				kPos := i - left
				if z[kPos] < right-i+1 {
					z[i] = z[kPos]
				} else {
					left = i
					for right < n && s[right-left] == s[right] {
						right++
					}

					z[i] = right - left
					right--
				}
			}
		}

		return z
	}

	z := zAlgorithm(word)

	// Check positions at multiples of k
	steps := 1

	pos := k
	for pos < n {
		// If remaining length matches prefix
		if z[pos] >= n-pos {
			return steps
		}

		pos += k
		steps++
	}

	return steps
}

// // Approach #1: Brute Force - Check each step
// // After each operation, remove first k chars and check if remaining matches prefix
// // Time: O(n^2 / k) where n = len(word)
// // Space: O(n) for string operations
// func minimumTimeToInitialState(word string, k int) int {
// 	n := len(word)
// 	steps := 1
// 	pos := k

// 	for pos < n {
// 		// Check if word[pos:] is a prefix of word
// 		if strings.HasPrefix(word, word[pos:]) {
// 			return steps
// 		}

// 		pos += k
// 		steps++
// 	}

// 	return steps
// }
