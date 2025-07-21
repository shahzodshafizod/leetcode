package stacks

import (
	"strings"
)

// https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/

// Approach: Greedy + Stack
// Time: O(n)
// Space: O(n)
func robotWithString(s string) string {
	var cnt [26]int
	for idx := range s {
		cnt[int(s[idx]-'a')]++
	}
	n := len(s)
	t := make([]int, n)
	size := 0
	var paper strings.Builder
	paper.Grow(n)
	minChar := 0
	for idx := range s {
		t[size] = int(s[idx] - 'a')
		size++
		cnt[t[size-1]]--
		for minChar < 25 && cnt[minChar] == 0 {
			minChar++
		}
		for size > 0 && t[size-1] <= minChar {
			paper.WriteByte(byte(t[size-1] + 'a'))
			size--
		}
	}
	return paper.String()
}

// // Time: O(n)
// // Space: O(n)
// func robotWithString(s string) string {
// 	var n = len(s)
// 	var right = make([]byte, n)
// 	right[n-1] = s[n-1]
// 	for idx := n - 2; idx >= 0; idx-- {
// 		right[idx] = min(s[idx], right[idx+1])
// 	}
// 	var t = make([]byte, n)
// 	var size = 0
// 	var paper strings.Builder
// 	paper.Grow(n)
// 	for idx := 0; idx < n; idx++ {
// 		for size > 0 && t[size-1] <= right[idx] {
// 			paper.WriteByte(t[size-1])
// 			size--
// 		}
// 		t[size] = s[idx]
// 		size++
// 	}
// 	for size--; size >= 0; size-- {
// 		paper.WriteByte(t[size])
// 	}
// 	return paper.String()
// }
