package strings

import (
	"slices"
)

// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/

func removeAnagrams(words []string) []string {
	res := []string{words[0]}
	curr := []byte(words[0])
	slices.Sort(curr)
	for i, n := 1, len(words); i < n; i++ {
		prev := curr
		curr = []byte(words[i])
		slices.Sort(curr)
		if string(prev) != string(curr) {
			res = append(res, words[i])
		}
	}
	return res
}
