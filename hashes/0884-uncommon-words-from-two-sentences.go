package hashes

import (
	"strings"
)

// https://leetcode.com/problems/uncommon-words-from-two-sentences/

func uncommonFromSentences(s1 string, s2 string) []string {
	allow := make(map[string]bool)
	deny := make(map[string]bool)
	for _, word := range strings.Split(s1+" "+s2, " ") {
		if deny[word] {
			continue
		}
		if allow[word] {
			deny[word] = true
			delete(allow, word)
			continue
		}
		allow[word] = true
	}
	uncommon := make([]string, 0, len(allow))
	for word := range allow {
		uncommon = append(uncommon, word)
	}
	return uncommon
}
