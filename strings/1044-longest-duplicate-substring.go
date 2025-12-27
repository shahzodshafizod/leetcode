package strings

// https://leetcode.com/problems/longest-duplicate-substring/

// Approach: Binary Search + Rolling Hash (Rabin-Karp Algorithm)
// Key insight: Binary search on the length of duplicate substring.
// For each length, use rolling hash to efficiently check if duplicate exists.
//
// Strategy:
// 1. Binary search on length: if length L has duplicate, try L+1
// 2. Rolling hash: compute hash for each substring of given length
// 3. Store hashes in map with indices to handle collisions
// 4. When hash collision found, verify actual strings match
//
// Time Complexity: O(n log n) - binary search O(log n) × rolling hash O(n)
// Space Complexity: O(n) for hash map
func longestDupSubstring(s string) string {
	// Helper function to search for duplicate substring of given length
	search := func(s string, length int) string {
		const (
			base = 26         // For lowercase letters a-z
			mod  = 1000000007 // Prime modulo to reduce collisions
		)

		n := len(s)

		// Compute base^(length-1) % mod
		power := 1
		for range length - 1 {
			power = (power * base) % mod
		}

		// Compute initial hash for first window
		hash := 0
		for i := range length {
			hash = (hash*base + int(s[i]-'a')) % mod
		}

		// Map hash to list of starting indices (to handle collisions)
		seen := map[int][]int{hash: {0}}

		// Slide the window and compute rolling hash
		for i := length; i < n; i++ {
			// Remove leftmost character
			hash = (hash - int(s[i-length]-'a')*power) % mod
			if hash < 0 {
				hash += mod
			}

			// Add rightmost character
			hash = (hash*base + int(s[i]-'a')) % mod

			// Check if this hash was seen before
			if indices, ok := seen[hash]; ok {
				current := s[i-length+1 : i+1]
				// Verify actual strings match (handle hash collisions)
				for _, idx := range indices {
					if s[idx:idx+length] == current {
						return current
					}
				}
				// Hash collision but strings don't match
				seen[hash] = append(seen[hash], i-length+1)
			} else {
				seen[hash] = []int{i - length + 1}
			}
		}

		return ""
	}

	n := len(s)
	left, right := 1, n-1
	result := ""

	// Binary search on substring length
	for left <= right {
		mid := left + (right-left)/2
		if dup := search(s, mid); dup != "" {
			result = dup
			left = mid + 1 // Try longer length
		} else {
			right = mid - 1 // Try shorter length
		}
	}

	return result
}
