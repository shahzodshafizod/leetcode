package maths

// https://leetcode.com/problems/water-bottles/

func numWaterBottles(numBottles int, numExchange int) int {
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
