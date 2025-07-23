package matrices

// https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	wealth := 0
	var total int
	for _, money := range accounts {
		total = 0
		for _, m := range money {
			total += m
		}
		wealth = max(wealth, total)
	}
	return wealth
}
