package maths

// https://leetcode.com/problems/calculate-money-in-leetcode-bank/

func totalMoney(n int) int {
	money, total := 0, 0

	for day := range n {
		if day%7 == 0 {
			money = day/7 + 1
		}

		total += money
		money++
	}

	return total
}
