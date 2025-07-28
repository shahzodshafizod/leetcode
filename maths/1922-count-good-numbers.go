package maths

// https://leetcode.com/problems/count-good-numbers/

// Approach: Binary Exponentiation
// Time: O(logn)
// Space: O(logn)
func countGoodNumbers(n int64) int {
	const MOD = int(1e9) + 7

	var pow func(x int, n int64) int

	pow = func(x int, n int64) int {
		if n == 0 {
			return 1
		}

		half := pow(x, n/2)
		if n&1 == 1 {
			return half * half * x % MOD
		}

		return half * half % MOD
	}
	evens := pow(5, n-n/2)
	primes := pow(4, n/2)

	return evens * primes % MOD
}
