package maths

// https://leetcode.com/problems/self-dividing-numbers/

// Approach 2: Optimized (Mathematical Digit Extraction)
// Idea: Extract digits mathematically without string conversion
// Time Complexity: O((right - left) * log(max_num))
//   - Same as brute force but more efficient constant factor
//
// Space Complexity: O(k), where k is the count of self-dividing numbers
func selfDividingNumbers(left int, right int) []int {
	isSelfDividing := func(num int) bool {
		original := num
		for num > 0 {
			digit := num % 10
			// If digit is 0 or doesn't divide original number
			if digit == 0 || original%digit != 0 {
				return false
			}

			num /= 10
		}

		return true
	}

	result := make([]int, 0)

	for num := left; num <= right; num++ {
		if isSelfDividing(num) {
			result = append(result, num)
		}
	}

	return result
}

// // Approach 1: Brute Force with String Conversion
// // Idea: Iterate through range, convert number to string to check each digit
// // Time Complexity: O((right - left) * log(max_num))
// //   - Iterate through (right - left + 1) numbers
// //   - For each number, convert to string and check digits (log(max_num) operations)
// //
// // Space Complexity: O(k), where k is the count of self-dividing numbers in result
// func selfDividingNumbers(left int, right int) []int {
// 	result := make([]int, 0)

// 	for num := left; num <= right; num++ {
// 		// Convert to string to easily access each digit
// 		numStr := strconv.Itoa(num)
// 		isSelfDividing := true

// 		for _, digitChar := range numStr {
// 			digit := int(digitChar - '0')
// 			// Check if digit is 0 or doesn't divide the number
// 			if digit == 0 || num%digit != 0 {
// 				isSelfDividing = false
// 				break
// 			}
// 		}

// 		if isSelfDividing {
// 			result = append(result, num)
// 		}
// 	}

// 	return result
// }
