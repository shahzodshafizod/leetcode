package greedy

import "sort"

// https://leetcode.com/problems/apple-redistribution-into-boxes/

// Approach: Greedy + Sorting
// Calculate total apples, sort boxes by capacity (largest first)
// Greedily select boxes until all apples fit
// N = len(apple), M = len(capacity)
// Time: O(N + M log M) - sum apples O(N), sort capacities O(M log M), iterate O(M)
// Space: O(1) - sort in-place (or O(M) if considering sort space)
func minimumBoxes(apple []int, capacity []int) int {
	// Calculate total apples to distribute
	totalApples := 0
	for _, a := range apple {
		totalApples += a
	}

	// Sort capacities in descending order (largest boxes first)
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	// Greedily select boxes
	boxesNeeded := 0
	currentCapacity := 0

	for _, boxCapacity := range capacity {
		currentCapacity += boxCapacity
		boxesNeeded++

		// If we can fit all apples, we're done
		if currentCapacity >= totalApples {
			return boxesNeeded
		}
	}

	// This shouldn't happen given problem constraints
	return boxesNeeded
}
