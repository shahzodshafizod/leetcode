package twopointers

import (
	"slices"
	"strings"
)

// https://leetcode.com/problems/reverse-words-in-a-string-iii/

// Approach: built-in functions
// Time: O(n)
// Space: O(n)
func reverseWords(s string) string {
	var words = strings.Split(s, " ")
	for idx := range words {
		word := []byte(words[idx])
		slices.Reverse(word)
		words[idx] = string(word)
	}
	return strings.Join(words, " ")
}

// // Approach: Two Pointers
// // Time: O(n)
// // Space: O(1)
// func reverseWords(s string) string {
// 	var b = []byte(s)
// 	var start, n = 0, len(s)
// 	for end := range s {
// 		if s[end] == ' ' || end == n-1 {
// 			var last = end
// 			if s[end] == ' ' {
// 				end--
// 			}
// 			for start < end {
// 				b[start], b[end] = b[end], b[start]
// 				start++
// 				end--
// 			}
// 			start = last + 1
// 		}
// 	}
// 	return string(b)
// }
