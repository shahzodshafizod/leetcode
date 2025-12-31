package strings

// https://leetcode.com/problems/to-lower-case/

// Approach #2: ASCII manipulation
// Manually convert each uppercase letter using ASCII math
// Uppercase letters: A-Z have ASCII codes 65-90
// Lowercase letters: a-z have ASCII codes 97-122
// Difference: 'a' - 'A' = 32
// N = length of string
// Time: O(N) - iterate through all characters
// Space: O(N) - create new byte slice/string
func toLowerCase(s string) string {
	result := []byte(s)
	for i := range result {
		// Check if character is uppercase (A-Z)
		if result[i] >= 'A' && result[i] <= 'Z' {
			// Convert to lowercase by adding 32 to ASCII value
			result[i] += 32
		}
	}

	return string(result)
}

// // Approach #1: Built-in function
// // Use Go's built-in strings.ToLower() function
// // N = length of string
// // Time: O(N) - iterate through all characters
// // Space: O(N) - create new string
// func toLowerCase(s string) string {
// 	return strings.ToLower(s)
// }
