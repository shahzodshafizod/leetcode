package maths

// https://leetcode.com/problems/ugly-number/

// Approach: Divide by 2, 3, 5 repeatedly
// Keep dividing by 2, 3, 5 while divisible.
// If result is 1, it's an ugly number.
// Time: O(log n) - divide operations
// Space: O(1) - constant space
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, factor := range []int{2, 3, 5} {
		for n%factor == 0 {
			n /= factor
		}
	}

	return n == 1
}
