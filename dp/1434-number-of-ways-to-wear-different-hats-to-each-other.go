package dp

// https://leetcode.com/problems/number-of-ways-to-wear-different-hats-to-each-other/

// Approach: Top-Down Dynamic Programming + Bitmasks
// Time: O(k⋅n⋅2^n), k=# of hats, n=# of people
// Space: O(k⋅2^n)
func numberWays(hats [][]int) int {
	const MOD int = 1e9 + 7
	var people = make(map[int][]int)
	for person, hts := range hats {
		for _, hat := range hts {
			people[hat-1] = append(people[hat-1], person)
		}
	}
	var n = len(hats)
	var done = (1 << n) - 1
	const MAX_HAT = 40
	var memo [MAX_HAT][]*int
	for idx := range memo {
		memo[idx] = make([]*int, done)
	}
	var dp func(hat, mask int) int
	dp = func(hat, mask int) int {
		if mask == done {
			return 1
		}
		if hat == MAX_HAT {
			return 0
		}
		if memo[hat][mask] != nil {
			return *memo[hat][mask]
		}
		// 1. skip
		var ways = dp(hat+1, mask)
		// 2. include
		for _, person := range people[hat] {
			person = 1 << person
			if mask&person == 0 {
				ways = (ways + dp(hat+1, mask|person)) % MOD
			}
		}
		memo[hat][mask] = &ways
		return ways
	}
	return dp(0, 0)
}
