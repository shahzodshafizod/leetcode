package maths

// https://leetcode.com/problems/total-characters-in-string-after-transformations-ii/

// Approach: Matrix Multiplication + Matrix Exponentiation By Squaring
// Time: O(n+logt×26^3)
// Space: O(26^2)
func lengthAfterTransformations(s string, t int, nums []int) int {
	const MOD int = 1e9 + 7
	multiply := func(a, b [26][26]int) [26][26]int {
		var c [26][26]int
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if a[i][j] != 0 {
					for k := 0; k < 26; k++ {
						c[i][k] = (c[i][k] + a[i][j]*b[j][k]) % MOD
					}
				}
			}
		}
		return c
	}
	powMatrix := func(base [26][26]int, exp int) [26][26]int {
		// identity matrix
		var res [26][26]int
		for i := 0; i < 26; i++ {
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
	for i := 0; i < 26; i++ {
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
		for j := 0; j < 26; j++ {
			rowSum = (rowSum + final[j][letter]) % MOD
		}
		res = (res + rowSum) % MOD
	}
	return res
}
