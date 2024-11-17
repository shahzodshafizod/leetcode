package maths

// https://leetcode.com/problems/add-digits/

// Follow up: Could you do it without any loop/recursion in O(1) runtime?
// Approach #4: Digital root (Time-Optimized)
// Time: O(1)
// Space: O(1)
func addDigits(num int) int {
	// A simple idea why digital root equals to mod 9 if we've got an ABCD number
	// ABCD = 1000A+100B+10*C+D = (A + B + C + D) + 9 * (111 * A + 11 * B + C)
	// this equals (mod 9) to A + B + C + D.

	// if num == 0 {
	// 	return 0
	// }
	// if num%9 == 0 {
	// 	return 9
	// }
	// return num % 9

	if num == 0 {
		return 0
	}
	return (num-1)%9 + 1
}

// // Approach #3: Digital root
// // Time: O(D), D=# of digits
// // Space: O(1)
// func addDigits(num int) int {
// 	for num > 9 {
// 		num = num/10 + num%10
// 	}
// 	return num
// }

// // Approach #2: Iterative
// // Time: O(D), D=# of digits
// // Space: O(1)
// func addDigits(num int) int {
// 	for num >= 10 {
// 		var sum = 0
// 		for num > 0 {
// 			sum += num % 10
// 			num /= 10
// 		}
// 		num = sum
// 	}
// 	return num
// }

// // Approach #1: Recursive
// // Time: O(D), D=# of digits
// // Space: O(D), for recursion stack
// func addDigits(num int) int {
// 	if num < 10 {
// 		return num
// 	}
// 	var sum = 0
// 	for num > 0 {
// 		sum += num % 10
// 		num /= 10
// 	}
// 	return addDigits(sum)
// }
