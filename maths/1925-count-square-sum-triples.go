package maths

// https://leetcode.com/problems/count-square-sum-triples/

// Approach 1: Brute Force - Try All Combinations
// Try all possible combinations of a, b, c
// Check if a² + b² = c²
// Time: O(n³) - too slow
// Space: O(1)

// Approach 2: Optimized - Two Loops with Set
// For each c, try all pairs (a, b) where a < c and b < c
// Check if a² + b² = c²
// Time: O(n²)
// Space: O(1)

// Approach 3: Hash Set Optimization
// For each c, iterate through a, calculate b² = c² - a²
// Check if b is a perfect square and b <= n
// Time: O(n²)
// Space: O(1)
func countTriples(n int) int {
	// Integer square root helper
	intSqrt := func(n int) int {
		if n < 2 {
			return n
		}

		// Binary search for square root
		left, right := 1, n/2
		for left <= right {
			mid := left + (right-left)/2
			square := mid * mid

			if square == n {
				return mid
			} else if square < n {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		return right
	}

	count := 0

	// Try all possible values of c (hypotenuse)
	for c := 1; c <= n; c++ {
		cSquared := c * c

		// Try all possible values of a
		for a := 1; a < c; a++ {
			aSquared := a * a

			// Calculate b² = c² - a²
			bSquared := cSquared - aSquared

			// Check if b² is a perfect square
			b := intSqrt(bSquared)
			if b*b == bSquared && b <= n && b > 0 {
				count++
			}
		}
	}

	return count
}

// // Alternative: Double loop checking all a and b
// func countTriples(n int) int {
// 	// Integer square root helper
// 	intSqrt := func(n int) int {
// 		if n < 2 {
// 			return n
// 		}

// 		// Binary search for square root
// 		left, right := 1, n/2
// 		for left <= right {
// 			mid := left + (right-left)/2
// 			square := mid * mid

// 			if square == n {
// 				return mid
// 			} else if square < n {
// 				left = mid + 1
// 			} else {
// 				right = mid - 1
// 			}
// 		}

// 		return right
// 	}

// 	count := 0

// 	for a := 1; a <= n; a++ {
// 		for b := 1; b <= n; b++ {
// 			sumSquares := a*a + b*b

// 			// Check if sum is a perfect square
// 			c := intSqrt(sumSquares)
// 			if c*c == sumSquares && c <= n {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }
