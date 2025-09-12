package strings

import "strings"

// https://leetcode.com/problems/vowels-game-in-a-string/

func doesAliceWin(s string) bool {
	// vowels := map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}

	// var count int
	// for _, c := range s {
	// 	count += vowels[c]
	// }

	// return count > 0
	return strings.ContainsAny(s, "aeiou")
}
