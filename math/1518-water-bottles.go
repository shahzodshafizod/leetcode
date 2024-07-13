package math

// https://leetcode.com/problems/water-bottles/

func numWaterBottles(numBottles int, numExchange int) int {
	var consumed = 0
	var empty = 0
	for numBottles > 0 {
		consumed += numBottles
		empty += numBottles
		numBottles = empty / numExchange
		empty %= numExchange
	}
	return consumed
}
