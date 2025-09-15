package hashes

import "strings"

// https://leetcode.com/problems/maximum-number-of-words-you-can-type/

func canBeTypedWords(text string, brokenLetters string) int {
	broken := make(map[rune]bool, 26)
	for _, c := range brokenLetters {
		broken[c] = true
	}

	var count int
	for _, word := range strings.Split(text, " ") {
		count++

		for _, c := range word {
			if broken[c] {
				count--

				break
			}
		}
	}

	return count
}
