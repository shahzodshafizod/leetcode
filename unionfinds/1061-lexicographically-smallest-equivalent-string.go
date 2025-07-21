package unionfinds

import (
	"strings"
)

// https://leetcode.com/problems/lexicographically-smallest-equivalent-string/

// Time: O(n+m)
// Space: O(1)
func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	var root [26]int
	for idx := 0; idx < 26; idx++ {
		root[idx] = idx
	}
	var find func(x int) int
	find = func(x int) int {
		if root[x] != x {
			root[x] = find(root[x])
		}
		return root[x]
	}
	n := len(s1)
	var x, y int
	for idx := 0; idx < n; idx++ {
		x = find(int(s1[idx] - 'a'))
		y = find(int(s2[idx] - 'a'))
		if x != y {
			root[max(x, y)] = min(x, y)
		}
	}
	var b strings.Builder
	for _, c := range baseStr {
		x = int(c - 'a')
		y = find(x)
		b.WriteByte(byte('a' + y))
	}
	return b.String()
}
