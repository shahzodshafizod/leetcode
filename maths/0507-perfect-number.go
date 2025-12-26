package maths

// https://leetcode.com/problems/perfect-number/

// Approach 1: Brute Force - Check All Divisors
// Iterate from 1 to n-1 and sum up all divisors
// Time: O(n) - too slow for large numbers
// Space: O(1)

// // Approach 2: Optimized - Check up to sqrt(n)
// // Key insight: divisors come in pairs (i, n/i)
// // Only need to check up to sqrt(n)
// // Time: O(sqrt(n))
// // Space: O(1)
// func checkPerfectNumber(num int) bool {
// 	// Perfect numbers must be greater than 1
// 	if num <= 1 {
// 		return false
// 	}

// 	sum := 1 // 1 is always a divisor

// 	// Check divisors up to sqrt(num)
// 	for i := 2; i*i <= num; i++ {
// 		if num%i == 0 {
// 			sum += i
// 			// Add the paired divisor if it's different
// 			if i*i != num {
// 				sum += num / i
// 			}
// 		}
// 	}

// 	return sum == num
// }

// Alternative: Early termination if sum exceeds num
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	sum := 1

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i != num {
				sum += num / i
			}
			// Early termination
			if sum > num {
				return false
			}
		}
	}

	return sum == num
}
