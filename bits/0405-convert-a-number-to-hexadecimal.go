package bits

// https://leetcode.com/problems/convert-a-number-to-hexadecimal/

// Approach 1: Brute Force - Manual Conversion
// Keep dividing by 16 and collecting remainders
// Handle negative numbers with two's complement
// Time: O(log16(n)) = O(8) for 32-bit = O(1)
// Space: O(8) = O(1)

// // Approach 2: Bit Manipulation (Optimal)
// // Extract 4 bits at a time from right to left
// // Each 4 bits represents one hex digit
// // Handle negative numbers using unsigned representation
// // Time: O(8) = O(1) - at most 8 hex digits for 32-bit int
// // Space: O(8) = O(1)
// func toHex(num int) string {
// 	if num == 0 {
// 		return "0"
// 	}

// 	// Convert to unsigned 32-bit representation
// 	// This handles negative numbers via two's complement
// 	unsigned := uint32(num)

// 	hexChars := "0123456789abcdef"
// 	result := ""

// 	// Extract 4 bits at a time (each hex digit)
// 	for unsigned > 0 {
// 		digit := unsigned & 0xF // Get last 4 bits
// 		result = string(hexChars[digit]) + result
// 		unsigned >>= 4 // Shift right by 4 bits
// 	}

// 	return result
// }

// Alternative implementation building result forward
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	unsigned := uint32(num)
	hexChars := "0123456789abcdef"
	digits := []byte{}

	for unsigned > 0 {
		digits = append(digits, hexChars[unsigned&0xF])
		unsigned >>= 4
	}

	// Reverse digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
