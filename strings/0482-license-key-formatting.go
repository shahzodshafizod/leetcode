package strings

import (
	"strings"
)

// https://leetcode.com/problems/license-key-formatting/

// Approach: Left to Right Traversal
// Time: O(n)
// Space: O(1)
func licenseKeyFormatting(s string, k int) string {
	n := len(s)

	count := (n - strings.Count(s, "-")) % k
	if count == 0 {
		count = k
	}

	var (
		sb  strings.Builder
		err error
	)

	for idx := range n {
		if s[idx] != '-' {
			if count == 0 {
				err = sb.WriteByte('-')
				_ = err

				count = k
			}

			err = sb.WriteByte(s[idx])
			_ = err

			count--
		}
	}

	return strings.ToUpper(sb.String())
}
