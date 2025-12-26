package maths

// https://leetcode.com/problems/base-7/

// Approach 1: Brute Force using built-in conversion
// Could use strconv.FormatInt(int64(num), 7) but problem says
// implement conversion yourself
// Time: O(log7(n))
// Space: O(log7(n))

// // Approach 2: Mathematical Conversion (Optimal)
// // Repeatedly divide by 7 and collect remainders
// // Build result string from remainders in reverse order
// // Time: O(log7(n)) - number of digits in base 7
// // Space: O(log7(n)) - for result string
// func convertToBase7(num int) string {
// 	// Handle zero case
// 	if num == 0 {
// 		return "0"
// 	}

// 	// Handle negative numbers
// 	isNegative := num < 0
// 	if isNegative {
// 		num = -num
// 	}

// 	result := ""

// 	// Convert to base 7
// 	for num > 0 {
// 		remainder := num % 7
// 		result = strconv.Itoa(remainder) + result
// 		num /= 7
// 	}

// 	if isNegative {
// 		result = "-" + result
// 	}

// 	return result
// }

// Alternative using string builder for efficiency
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := num < 0
	if isNegative {
		num = -num
	}

	digits := []byte{}

	for num > 0 {
		digits = append(digits, byte('0'+num%7))
		num /= 7
	}

	// Reverse digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if isNegative {
		return "-" + string(digits)
	}

	return string(digits)
}
