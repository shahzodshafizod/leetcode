package maths

// https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/

// Time: O(N)
// Space: O(1)
func numberOfWays(corridor string) int {
	const MOD int = 1e9 + 7
	seats, plants := 0, 0
	ways := 1
	for _, c := range corridor {
		if c == 'S' {
			seats++
		} else if seats == 2 {
			plants++
		}
		if seats == 3 {
			ways = (ways * (plants + 1)) % MOD
			seats, plants = 1, 0
		}
	}
	if seats != 2 {
		ways = 0
	}
	return ways
}
