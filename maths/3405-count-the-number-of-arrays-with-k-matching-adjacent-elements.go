package maths

// https://leetcode.com/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/

/*
comb(n, k) = n! / (k! * (n-k)!)

return fact[n] / (fact[k] * fact[n-k]) % MOD

Here you're doing normal division (/), but in modular arithmetic, you cannot divide normally.
In modular arithmetic, division must be replaced by multiplication with modular inverse!

Why?
Because (a / b) % MOD != (a % MOD) / (b % MOD) — that's mathematically wrong under modulo.

Instead, the correct formula is:
fact[n] * modInverse(fact[k] * fact[n-k] % MOD, MOD) % MOD

You must multiply by modular inverse of the denominator under modulo.
*/

// Approach: Combinatorics
// Time: O(Min(K,N−K))
// Space: O(1)
func countGoodArrays(n int, m int, k int) int {
	const mod int = 1e9 + 7

	pow := func(base, exp int) int {
		result := 1

		for exp > 0 {
			if exp&1 == 1 {
				result = result * base % mod
			}

			base = base * base % mod
			exp >>= 1
		}

		return result
	}
	comb := func(n int, k int) int {
		if k > n/2 {
			k = n - k
		}

		num := 1 // numerator
		den := 1 // denominator

		for i := 0; i < k; i++ {
			num = (num * (n - i)) % mod
			den = (den * (i + 1)) % mod
		}

		den = pow(den, mod-2)

		return num * den % mod
	}

	return m * pow(m-1, n-1-k) % mod * comb(n-1, k) % mod
}
