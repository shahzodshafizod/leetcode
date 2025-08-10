package maths

// https://leetcode.com/problems/reordered-power-of-2/

func reorderedPowerOf2(n int) bool {
	var cnt [10]int
	for ; n > 0; n /= 10 {
		cnt[n%10]++
	}

	var backtrack func(int) bool

	backtrack = func(n int) bool {
		d := n % 10
		if n == 0 {
			total := 0
			for _, c := range cnt {
				total += c
			}

			return total == 0
		}

		if cnt[d] == 0 {
			return false
		}

		cnt[d]--
		res := backtrack(n / 10)
		cnt[d]++

		return res
	}
	for i := range 32 {
		if backtrack(1 << i) {
			return true
		}
	}

	return false
}
