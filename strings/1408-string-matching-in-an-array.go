package strings

import "sort"

// https://leetcode.com/problems/string-matching-in-an-array/

// Approach: KMP (Knutt-Morris-Pritt)
// Time: O(nnm), n=len(words), m=len(words[i])
// Space: O(m)
func stringMatching(words []string) []string {
	getLPS := func(word string) []int {
		m := len(word)
		lps := make([]int, m)

		p, s := 0, 1
		for s < m {
			if word[s] == word[p] {
				p++
				lps[s] = p
				s++
			} else if p == 0 {
				s++
			} else {
				p = lps[p-1]
			}
		}

		return lps
	}

	sort.Slice(words, func(i int, j int) bool { return len(words[i]) < len(words[j]) })
	n := len(words)
	check := func(curr string, start int) bool {
		m := len(curr)
		lps := getLPS(curr)

		var cid, nid, nlen int

		for idx := start; idx < n; idx++ {
			next := words[idx]

			nlen = len(next)
			if nlen == m {
				continue
			}

			cid, nid = 0, 0
			for cid < m && nid < nlen {
				if curr[cid] == next[nid] {
					cid++
					nid++
				} else if cid == 0 {
					nid++
				} else {
					cid = lps[cid-1]
				}
			}

			if cid == m {
				return true
			}
		}

		return false
	}
	result := make([]string, 0)

	for idx, word := range words {
		if check(word, idx+1) {
			result = append(result, word)
		}
	}

	return result
}

// // Approach: Trie
// // Time: O(nmm), n=len(words), m=len(words[i])
// // Space: O(nm)
// func stringMatching(words []string) []string {
// 	type TrieNode struct {
// 		count int
// 		children map[rune]*TrieNode
// 	}
// 	var NewTrieNode = func() *TrieNode {
// 		return &TrieNode{children: make(map[rune]*TrieNode)}
// 	}
// 	var root = NewTrieNode()
// 	var n = len(words)
// 	var nodes = make([]*TrieNode, n)
// 	var m int
// 	for idx, word := range words {
// 		m = len(word)
// 		for j := 0; j < m; j++ {
// 			curr := root
// 			for _, c := range word[j:] {
// 				if curr.children[c] == nil {
// 					curr.children[c] = NewTrieNode()
// 				}
// 				curr = curr.children[c]
// 				curr.count++
// 			}
// 			if nodes[idx] == nil {
// 				nodes[idx] = curr
// 			}
// 		}
// 	}
// 	var result = make([]string, 0)
// 	for idx := range nodes {
// 		if nodes[idx].count >= 2 {
// 			result = append(result, words[idx])
// 		}
// 	}
// 	return result
// }
