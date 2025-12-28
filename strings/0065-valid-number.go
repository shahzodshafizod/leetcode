package strings

// https://leetcode.com/problems/valid-number/

// Approach: Optimized - State Machine / Character-by-character parsing
// Parse string while tracking: signs, digits, decimal point, exponent
// Valid number = (integer OR decimal) + optional exponent
// Time: O(N) - single pass through string
// Space: O(1) - only boolean flags
func isNumber(s string) bool {
	seenDigit := false
	seenExponent := false
	seenDot := false

	for i, char := range s {
		if char >= '0' && char <= '9' {
			seenDigit = true
		} else if char == '+' || char == '-' {
			// Sign is valid at start OR right after 'e'/'E'
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else if char == '.' {
			// Dot is invalid after exponent or another dot
			if seenExponent || seenDot {
				return false
			}

			seenDot = true
		} else if char == 'e' || char == 'E' {
			// Exponent requires at least one digit before it
			// and cannot appear twice
			if seenExponent || !seenDigit {
				return false
			}

			seenExponent = true
			seenDigit = false // Must have digit after exponent
		} else {
			// Invalid character
			return false
		}
	}

	return seenDigit
}
