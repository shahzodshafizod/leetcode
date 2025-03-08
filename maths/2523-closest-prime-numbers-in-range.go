package maths

import "math"

// https://leetcode.com/problems/closest-prime-numbers-in-range/

// Approach: Sieve of Eratosthenes
// Time: O(N∗Log(Log(N)))
// Space: O(N)
func closestPrimes(left int, right int) []int {
	var nonprime = make([]bool, right+1)
	nonprime[0], nonprime[1] = true, true
	for num, sqrt := 2, int(math.Sqrt(float64(right))); num <= sqrt; num++ {
		if !nonprime[num] {
			for factor := num + num; factor <= right; factor += num {
				nonprime[factor] = true
			}
		}
	}
	var ans = []int{-1, -1}
	var diff = right - left + 1
	var prev = -1
	for curr := left; curr <= right; curr++ {
		if nonprime[curr] {
			continue
		}
		if prev != -1 && curr-prev < diff {
			diff = curr - prev
			ans = []int{prev, curr}
		}
		prev = curr
	}
	return ans
}
