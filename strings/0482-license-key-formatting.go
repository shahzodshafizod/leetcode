package strings

import (
	"strings"
)

// https://leetcode.com/problems/license-key-formatting/

// Approach: Left to Right Traversal
// Time: O(n)
// Space: O(1)
func licenseKeyFormatting(s string, k int) string {
	var n = len(s)
	var count = (n - strings.Count(s, "-")) % k
	if count == 0 {
		count = k
	}
	var sb strings.Builder
	for idx := 0; idx < n; idx++ {
		if s[idx] != '-' {
			if count == 0 {
				sb.WriteByte('-')
				count = k
			}
			sb.WriteByte(s[idx])
			count--
		}
	}
	return strings.ToUpper(sb.String())
}
