package hashes

import (
	"strings"
)

// https://leetcode.com/problems/word-pattern/

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	n := len(words)
	if n != len(pattern) {
		return false
	}

	codes := make(map[string]byte)

	var used [26]bool

	var letter byte

	var ok bool

	var code int

	for idx := range n {
		if letter, ok = codes[words[idx]]; !ok {
			code = int(pattern[idx] - 'a')
			if used[code] {
				return false
			}

			used[code] = true
			codes[words[idx]] = pattern[idx]
		} else if letter != pattern[idx] {
			return false
		}
	}

	return true
}
