package bits

// https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/

// Approach: Gray Code
// Time: O(logn)
// Space: O(1)
func minimumOneBitOperations(n int) int {
	operations := 0
	for n > 0 {
		operations ^= n
		n /= 2
	}
	return operations
}

// // Approach: Iterative
// // Time: O(O(log^2(n)))
// // Space: O(1)
// func minimumOneBitOperations(n int) int {
// 	var mask, power2 = 1, 1
// 	var operations = 0
// 	for power2 <= n {
// 		power2 *= 2
// 		if mask&n != 0 {
// 			operations = power2 - 1 - operations
// 		}
// 		mask <<= 1
// 	}
// 	return operations
// }

// // Approach: Recursive
// // Time: O(O(log^2(n)))
// // Space: O(logn)
// func minimumOneBitOperations(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	var power2 = 1
// 	for power2 <= n {
// 		power2 *= 2
// 	}
// 	return power2 - 1 - minimumOneBitOperations(power2/2^n)
// }
