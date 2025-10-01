package maths

// https://leetcode.com/problems/water-bottles/

// Approach: Simulation
// Time: O(logm(n))
// Space: O(1)
func numWaterBottles(numBottles int, numExchange int) int {
	// return numBottles + (numBottles-1)/(numExchange-1)
	consumed := 0
	empty := 0

	for numBottles > 0 {
		consumed += numBottles
		empty += numBottles
		numBottles = empty / numExchange
		empty %= numExchange
	}

	return consumed
}
