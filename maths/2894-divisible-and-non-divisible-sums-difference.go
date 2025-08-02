package maths

/*
num2 = 1m + 2m + 3m + 4m
num2 = (1 + 2 + 3 + 4) * m
num2 = (k * (k+1) // 2) * m
num1 = n * (n+1) // 2 - num2
res = num1 - num2
res = n * (n+1) // 2 - num2 - num2
res = n * (n+1) // 2 - 2 * num2
res = n * (n+1) // 2 - [2] * (k * (k+1) // [2]) * m
res = n * (n+1) // 2 - k * (k+1) * m
*/

// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/

func differenceOfSums(n int, m int) int {
	k := n / m

	return n*(n+1)/2 - k*(k+1)*m
}
