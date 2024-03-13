package math

import "math"

/*
n:6
0+1+2+3+4+5+6
+++++++++++++
6+5+4+3+2+1+0
=============
6+6+6+6+6+6+6

sum(n)=n*(n+1)/2
*/

// https://leetcode.com/problems/find-the-pivot-integer/

// time: O(1)
// space: O(1)
func pivotInteger(n int) int {
	var sum = int(float32(n) / 2 * float32(n+1))
	var pivot = int(math.Sqrt(float64(sum)))
	if pivot*pivot == sum {
		return pivot
	}
	return -1
}

// // binary search
// // time: O(log n)
// func pivotInteger(n int) int {
// 	var sum = int(float32(n) / 2 * float32(n+1))
// 	var mid int
// 	var left, right = 0, n
// 	for left < right {
// 		mid = (left + right) / 2
// 		if mid*mid < sum {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}
// 	if left*left == sum {
// 		return left
// 	}
// 	return -1
// }
