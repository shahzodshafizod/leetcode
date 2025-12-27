package hashes

// https://leetcode.com/problems/distribute-candies/

// Approach #2: Hash Set (Optimal)
// Time: O(n) - single pass to build map
// Space: O(n) - map stores unique candy types
func distributeCandies(candyType []int) int {
	// Count unique candy types using map/set
	uniqueTypes := make(map[int]bool)
	for _, candy := range candyType {
		uniqueTypes[candy] = true
	}

	// Alice can eat at most n/2 candies
	maxEat := len(candyType) / 2

	// She wants to maximize variety, so return minimum of:
	// - number of unique types available
	// - maximum number she can eat (n/2)
	uniqueCount := len(uniqueTypes)
	if uniqueCount < maxEat {
		return uniqueCount
	}

	return maxEat
}

// // Approach #1: Brute-Force (Count unique types manually)
// // Time: O(n^2) - checking if each candy type exists in unique list
// // Space: O(n) - storing unique types
// func distributeCandies(candyType []int) int {
// 	// Find all unique candy types manually
// 	uniqueTypes := make([]int, 0)
// 	for _, candy := range candyType {
// 		isUnique := true
// 		for _, uniqueCandy := range uniqueTypes {
// 			if candy == uniqueCandy {
// 				isUnique = false
// 				break
// 			}
// 		}
// 		if isUnique {
// 			uniqueTypes = append(uniqueTypes, candy)
// 		}
// 	}

// 	// Alice can eat at most n/2 candies
// 	maxEat := len(candyType) / 2

// 	// Return minimum of unique types and max she can eat
// 	if len(uniqueTypes) < maxEat {
// 		return len(uniqueTypes)
// 	}
// 	return maxEat
// }
