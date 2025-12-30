package hashes

// https://leetcode.com/problems/jewels-and-stones/

// Approach: Optimized using Map (Hash Table)
// Convert jewels to a map for O(1) lookup time
// Time: O(n + m) where n is length of jewels, m is length of stones
// Space: O(n) where n is length of jewels (for the map)
func numJewelsInStones(jewels string, stones string) int {
	// Create a map to store jewel types for O(1) lookup
	jewelSet := make(map[rune]bool)

	for _, jewel := range jewels {
		jewelSet[jewel] = true
	}

	// Count how many stones are jewels
	count := 0

	for _, stone := range stones {
		if jewelSet[stone] {
			count++
		}
	}

	return count
}

// // Approach 1: Brute Force
// // For each stone, check if it exists in jewels string
// // Time: O(n * m) where n is length of stones, m is length of jewels
// // Space: O(1) - no extra data structures
// func numJewelsInStones(jewels string, stones string) int {
// 	count := 0

// 	// For each stone, check if it's a jewel
// 	for _, stone := range stones {
// 		// Linear search through jewels
// 		for _, jewel := range jewels {
// 			if stone == jewel {
// 				count++

// 				break
// 			}
// 		}
// 	}

// 	return count
// }
