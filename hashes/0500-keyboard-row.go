package hashes

import (
	"strings"
	"unicode"
)

// https://leetcode.com/problems/keyboard-row/

func findWords(words []string) []string {
	codes := make(map[rune]int)

	for idx, word := range [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"} {
		for _, c := range word {
			codes[c] = idx
		}
	}

	var res []string

	var code int

	var ok bool

	for _, word := range words {
		code = codes[unicode.ToLower(rune(word[0]))]
		ok = true

		for _, c := range strings.ToLower(word) {
			if codes[unicode.ToLower(c)] != code {
				ok = false
				break
			}
		}

		if ok {
			res = append(res, word)
		}
	}

	return res
}
