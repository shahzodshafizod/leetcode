package tries

import "strings"

// https://leetcode.com/problems/counting-words-with-a-given-prefix/

// Approach: Built-In Methods
// Time: O(m*n), m=len(pref), n=len(words)
// Space: O(1)
func prefixCount(words []string, pref string) int {
	var count = 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}
	return count
}

// // Approach: Brute-Force
// // Time: O(m*n), m=len(pref), n=len(words)
// // Space: O(1)
// func prefixCount(words []string, pref string) int {
// 	var plen = len(pref)
// 	var count = 0
// 	var pid, wid, wlen int
// 	for _, word := range words {
// 		pid, wid, wlen = 0, 0, len(word)
// 		for pid < plen && wid < wlen {
// 			if wlen-wid < plen-pid || word[wid] != pref[pid] {
// 				break
// 			}
// 			wid++
// 			pid++
// 		}
// 		if pid == plen {
// 			count += 1
// 		}
// 	}
// 	return count
// }

// // Approach: Tries
// // Time: O(m + n*l), m=len(pref), n=len(words), l=len(words[i])
// // Space: O(n*l)
// func prefixCount(words []string, pref string) int {
// 	type TrieNode struct {
// 		count int
// 		children map[rune]*TrieNode
// 	}
// 	var NewTrieNode = func() *TrieNode {
// 		return &TrieNode{children: make(map[rune]*TrieNode)}
// 	}
// 	var root = NewTrieNode()
// 	var curr *TrieNode
// 	for _, word := range words {
// 		curr = root
// 		for _, c := range word {
// 			if curr.children[c] == nil {
// 				curr.children[c] = NewTrieNode()
// 			}
// 			curr = curr.children[c]
// 			curr.count++
// 		}
// 	}
// 	var count = 0
// 	curr = root
// 	for _, c := range pref {
// 		if curr.children[c] == nil {
// 			count = 0
// 			break
// 		}
// 		curr = curr.children[c]
// 		count = curr.count
// 	}
// 	return count
// }
