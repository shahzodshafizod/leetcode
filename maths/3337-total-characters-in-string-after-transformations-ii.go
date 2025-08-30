package maths

// https://leetcode.com/problems/total-characters-in-string-after-transformations-ii/

// Approach: Matrix Multiplication + Matrix Exponentiation By Squaring
// Time: O(n+logt×26^3)
// Space: O(26^2)
func lengthAfterTransformations(s string, t int, nums []int) int {
	const mod int = 1e9 + 7

	multiply := func(a, b [26][26]int) [26][26]int {
		var c [26][26]int

		for i := range 26 {
			for j := range 26 {
				if a[i][j] != 0 {
					for k := range 26 {
						c[i][k] = (c[i][k] + a[i][j]*b[j][k]) % mod
					}
				}
			}
		}

		return c
	}
	powMatrix := func(base [26][26]int, exp int) [26][26]int {
		// identity matrix
		var res [26][26]int
		for i := range 26 {
			res[i][i] = 1
		}

		for ; exp > 0; exp >>= 1 {
			if exp&1 == 1 {
				res = multiply(res, base)
			}

			base = multiply(base, base)
		}

		return res
	}

	var transformation [26][26]int

	for i := range 26 {
		for j := 1; j <= nums[i]; j++ {
			transformation[(i+j)%26][i] = 1
		}
	}

	final := powMatrix(transformation, t)
	res := 0

	var letter, rowSum int
	for _, c := range s {
		letter = int(c - 'a')
		rowSum = 0

		for j := range 26 {
			rowSum = (rowSum + final[j][letter]) % mod
		}

		res = (res + rowSum) % mod
	}

	return res
}
