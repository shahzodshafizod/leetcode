package hashes

// https://leetcode.com/problems/minimum-index-sum-of-two-lists/

// Approach #2: Hash Map (Optimized)
// Time: O(n + m) - single pass through both lists
// Space: O(n) - hash map for list1
func findRestaurant(list1 []string, list2 []string) []string {
	// Build hash map for list1: string -> index
	indexMap := make(map[string]int)
	for i, restaurant := range list1 {
		indexMap[restaurant] = i
	}

	minSum := len(list1) + len(list2)

	var result []string

	// Iterate through list2 and find common strings
	for j, restaurant := range list2 {
		if i, found := indexMap[restaurant]; found {
			indexSum := i + j

			if indexSum < minSum {
				// Found new minimum, reset result
				minSum = indexSum
				result = []string{restaurant}
			} else if indexSum == minSum {
				// Same minimum, add to result
				result = append(result, restaurant)
			}
		}
	}

	return result
}

// // Approach #1: Brute-Force (Check all pairs)
// // Time: O(n * m) - two nested loops
// // Space: O(1) - no extra space besides output
// func findRestaurant(list1 []string, list2 []string) []string {
// 	minSum := len(list1) + len(list2)
// 	var result []string

// 	// Check all pairs
// 	for i, str1 := range list1 {
// 		for j, str2 := range list2 {
// 			if str1 == str2 {
// 				indexSum := i + j
// 				if indexSum < minSum {
// 					minSum = indexSum
// 					result = []string{str1}
// 				} else if indexSum == minSum {
// 					result = append(result, str1)
// 				}
// 			}
// 		}
// 	}

// 	return result
// }
