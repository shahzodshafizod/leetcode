package strings

import "unicode"

// https://leetcode.com/problems/capitalize-the-title/

// Time: O(N)
// Space: O(N)
func capitalizeTitle(title string) string {
	result := []rune(title)
	start, n := 0, len(result)

	for idx := 0; idx <= n; idx++ {
		if idx == n || result[idx] == ' ' {
			if idx-start > 2 {
				result[start] = unicode.ToUpper(result[start])
			}

			start = idx + 1
		} else {
			result[idx] = unicode.ToLower(result[idx])
		}
	}

	return string(result)
}
