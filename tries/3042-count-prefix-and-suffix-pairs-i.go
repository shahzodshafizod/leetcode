package tries

import (
	"math/bits"
)

// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/

func countPrefixSuffixPairs(words []string) int {
	type TrieNode struct {
		prefix   int64
		suffix   int64
		children map[rune]*TrieNode
	}
	NewTrieNode := func() *TrieNode {
		return &TrieNode{children: make(map[rune]*TrieNode)}
	}
	root := NewTrieNode()
	var prefix, suffix bool
	var c rune
	var currp, currs *TrieNode
	count := 0
	for wid := len(words) - 1; wid >= 0; wid-- {
		word := words[wid]
		prefix = true
		currp = root
		for _, c = range word {
			if currp.children[c] == nil {
				prefix = false
				currp.children[c] = NewTrieNode()
			}
			currp = currp.children[c]
			currp.prefix |= 1 << wid
		}
		suffix = true
		currs = root
		for idx := len(word) - 1; idx >= 0; idx-- {
			c = rune(word[idx])
			if currs.children[c] == nil {
				suffix = false
				currs.children[c] = NewTrieNode()
			}
			currs = currs.children[c]
			currs.suffix |= 1 << wid
		}
		if prefix && suffix {
			count += bits.OnesCount(uint(currp.prefix&currs.suffix)) - 1
		}
	}
	return count
}

// func countPrefixSuffixPairs(words []string) int {
// 	var count = 0
// 	for i, n := 0, len(words); i < n; i++ {
// 		for j := i+1; j < n; j++ {
// 			if len(words[i]) <= len(words[j]) && strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
