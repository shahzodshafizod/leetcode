package maths

// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-ii/

// Approach: Pascal's Triangle + Modular Arithmetic
// The operation reduces string length by 1 each time: n → n-1 → ... → 2
// Final two digits are weighted sums with coefficients from Pascal's triangle row (n-2)
// We need to check if sum(C(n-2, i) * (s[i] - s[i+1])) ≡ 0 (mod 10)
// To compute C(n-2, k) mod 10 efficiently, factor out powers of 2 and 5
// Time: O(n), Space: O(n)
func hasSameDigits(s string) bool {
	n := len(s)
	if n == 2 {
		return s[0] == s[1]
	}

	// Precompute factorials as (twos, fives, other) tuples
	// where x! = 2^twos * 5^fives * other, with gcd(other, 10) = 1
	type triple struct{ twos, fives, other int }

	factorize := func(x int) triple {
		twos := 0

		for x%2 == 0 {
			x /= 2
			twos++
		}

		fives := 0

		for x%5 == 0 {
			x /= 5
			fives++
		}

		return triple{twos, fives, x % 10}
	}

	facts := []triple{{0, 0, 1}} // 0! = 1
	for i := 1; i < n-1; i++ {
		prev := facts[len(facts)-1]
		curr := factorize(i)
		facts = append(facts, triple{
			twos:  prev.twos + curr.twos,
			fives: prev.fives + curr.fives,
			other: (prev.other * curr.other) % 10,
		})
	}

	// Modular inverse mod 10 (only exists for 1, 3, 7, 9)
	inv := map[int]int{1: 1, 3: 7, 7: 3, 9: 9}

	// Compute C(n-2, k) mod 10
	comb := func(nVal, k int) int {
		if k < 0 || k > nVal {
			return 0
		}
		// C(n, k) = n! / (k! * (n-k)!)
		num := facts[nVal]
		denK := facts[k]
		denNK := facts[nVal-k]

		// Combine denominators
		denTwos := denK.twos + denNK.twos
		denFives := denK.fives + denNK.fives
		denOther := (denK.other * denNK.other) % 10

		// Compute result
		twos := num.twos - denTwos
		fives := num.fives - denFives
		other := (num.other * inv[denOther]) % 10

		// Convert back to single digit
		if twos >= 1 && fives >= 1 {
			return 0
		}

		if fives >= 1 {
			return 5
		}

		if twos == 0 {
			return other
		}
		// Powers of 2 mod 10 cycle: 2, 4, 8, 6, 2, 4, 8, 6, ...
		pow2Cycle := []int{6, 2, 4, 8}

		return (other * pow2Cycle[twos%4]) % 10
	}

	// Check if weighted sum is 0 mod 10
	total := 0

	for i := range n - 1 {
		coeff := comb(n-2, i)
		diff := int(s[i]-'0') - int(s[i+1]-'0')
		total = (total + coeff*diff + 100) % 10 // Add 100 to handle negative
	}

	return total == 0
}
