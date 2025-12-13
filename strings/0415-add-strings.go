package strings

// https://leetcode.com/problems/add-strings/

// Approach #2: Simulation with Two Pointers
// Time: O(max(n, m)) where n, m are lengths of num1, num2
// Space: O(max(n, m)) for result string
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	result := make([]byte, 0)

	for i >= 0 || j >= 0 || carry > 0 {
		// Get current digits (0 if index out of bounds)
		digit1, digit2 := 0, 0
		if i >= 0 {
			digit1 = int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			digit2 = int(num2[j] - '0')
			j--
		}

		// Calculate sum with carry
		total := digit1 + digit2 + carry
		carry = total / 10
		result = append(result, byte(total%10)+'0')
	}

	// Result is built backwards, so reverse it
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return string(result)
}

// // Approach #1: Brute Force - Convert to integers and add
// // Time: O(n + m) - string to int conversion
// // Space: O(1) - excluding output
// func addStrings(num1 string, num2 string) string {
// 	n1, _ := strconv.Atoi(num1)
// 	n2, _ := strconv.Atoi(num2)

// 	return strconv.Itoa(n1 + n2)
// }
