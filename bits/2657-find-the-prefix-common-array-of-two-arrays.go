package bits

import "math/bits"

// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/

// Approach: Bit Manipulation
// Time: O(n)
// Space: O(1)
func findThePrefixCommonArray(A []int, B []int) []int {
	var amask, bmask int64 = 0, 0

	C := make([]int, len(A))

	for idx := range A {
		amask |= 1 << A[idx]
		bmask |= 1 << B[idx]
		C[idx] = bits.OnesCount(uint(amask & bmask))
	}

	return C
}

// // Approach: Counting
// // Time: O(n)
// // Space: O(n)
// func findThePrefixCommonArray(A []int, B []int) []int {
// 	var n = len(A)
// 	var C = make([]int, n)
// 	var count = make([]int, n+1)
// 	var answer = 0
// 	for idx := range A {
// 		count[A[idx]]++
// 		if count[A[idx]] == 2 {
// 			answer++
// 		}
// 		count[B[idx]]++
// 		if count[B[idx]] == 2 {
// 			answer++
// 		}
// 		C[idx] = answer
// 	}
// 	return C
// }
