package strings

import (
	"unicode"
)

// https://leetcode.com/problems/detect-capital/

func detectCapitalUse(word string) bool {
	var capitals = 0
	var first = false
	for idx, c := range word {
		if unicode.IsUpper(c) {
			capitals++
			first = first || idx == 0
		}
	}
	return capitals == 0 || capitals == len(word) || capitals == 1 && first
}
