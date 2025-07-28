package strings

import "strings"

// https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/

func divideString(s string, k int, fill byte) []string {
	var parts []string
	for start, n := 0, len(s); start < n; start += k {
		parts = append(parts, s[start:min(start+k, n)])
	}

	if n := len(parts); len(parts[n-1]) < k {
		parts[n-1] += strings.Repeat(string(fill), k-len(parts[n-1]))
	}

	return parts
}
