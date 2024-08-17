package maths

import "math"

// https://leetcode.com/problems/sum-of-square-numbers/

// Approach #4: Two Pointers
// time: O(sqrt(c))
// space: O(1)
func judgeSquareSum(c int) bool {
	var a, b = 0, int(math.Sqrt(float64(c)))
	var sum int
	for a <= b {
		sum = a*a + b*b
		if sum < c {
			a++
		} else if sum > c {
			b--
		} else {
			return true
		}
	}
	return false
}

// // Approach #3: Binary Search
// // time: O(n log n)
// // space: O(1)
// func judgeSquareSum(c int) bool {
// 	var limit = int(math.Sqrt(float64(c)))
// 	var left, right, b, sum int
// 	for a := 1; a <= limit; a++ {
// 		left, right = a+1, limit
// 		for left <= right {
// 			b = (left + right) / 2
// 			sum = a*a + b*b
// 			if sum > c {
// 				right = b - 1
// 			} else if sum < c {
// 				left = b + 1
// 			} else {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // Approach #2: Hash Set
// // time: O(c)
// // space: O(c)
// func judgeSquareSum(c int) bool {
// 	var len = int(math.Sqrt(float64(c))) + 1
// 	var square = make(map[int]bool)
// 	for a := 0; a <= len; a++ {
// 		square[a*a] = true
// 	}
// 	for b := 0; b <= len; b++ {
// 		if square[c-b*b] {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Approach #1: Brute Force
// // time: O(c^2)
// // space: O(1)
// func judgeSquareSum(c int) bool {
// 	var limit = int(math.Sqrt(float64(c))) + 1
// 	var aa int
// 	for a := 0; a <= limit; a++ {
// 		aa = a * a
// 		for b := 0; b <= limit; b++ {
// 			if aa+b*b == c {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
