package maths

// https://leetcode.com/problems/closest-prime-numbers-in-range/

// Approach: Sieve of Eratosthenes
// Time: O(N∗Log(Log(N)))
// Space: O(N)
func closestPrimes(left int, right int) []int {
	var sieve = make([]bool, right+1)
	var primes []int
	for num := 2; num <= right; num++ {
		if !sieve[num] {
			if num >= left {
				primes = append(primes, num)
			}
			for factor := num * num; factor <= right; factor += num {
				sieve[factor] = true
			}
		}
	}
	var ans = []int{-1, -1}
	var diff = right - left + 1
	var prev = -1
	for _, curr := range primes {
		if prev != -1 && curr-prev < diff {
			diff = curr - prev
			ans = []int{prev, curr}
		}
		prev = curr
	}
	return ans
}
