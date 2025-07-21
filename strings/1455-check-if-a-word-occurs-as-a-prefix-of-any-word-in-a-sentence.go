package strings

// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/

// Approach #3: Two Pointers
// Time: O(n+m), n=len(sentence), m=len(searchWord)
// Space: O(1)
func isPrefixOfWord(sentence string, searchWord string) int {
	wordIndex := 0
	idx, m := 0, len(searchWord)
	block := false
	for _, c := range sentence {
		if c == ' ' {
			wordIndex++
			idx = 0
			block = false
			continue
		} else if block {
			continue
		} else if searchWord[idx] != byte(c) {
			block = true
			continue
		}
		idx++
		if idx == m {
			return wordIndex + 1
		}
	}
	return -1
}

// // Approach #2: Brute-Force
// // Time: O(n), n=len(sentence), m=len(searchWord)
// // Space: O(n)
// func isPrefixOfWord(sentence string, searchWord string) int {
// 	var m = len(searchWord)
// 	var index = -1
// 	for idx, word := range strings.Split(sentence, " ") {
// 		if len(word) < m {
// 			continue
// 		}
// 		index = -1
// 		for j := 0; j < m; j++ {
// 			if searchWord[j] != word[j] {
// 				index = -1
// 				break
// 			}
// 			index = idx + 1
// 		}
// 		if index != -1 {
// 			break
// 		}
// 	}
// 	return index
// }

// // Approach #1: Trie
// // Time: O(n+m), n=len(sentence), m=len(searchWord)
// // Space: O(n)
// func isPrefixOfWord(sentence string, searchWord string) int {
// 	type TrieNode struct {
// 		children map[rune]*TrieNode
// 		index    int
// 	}
// 	var newTrieNode = func(index int) *TrieNode {
// 		return &TrieNode{
// 			children: make(map[rune]*TrieNode),
// 			index:    index,
// 		}
// 	}
// 	var root = newTrieNode(-1)
// 	var curr *TrieNode
// 	for idx, word := range strings.Split(sentence, " ") {
// 		var curr = root
// 		for _, c := range word {
// 			if curr.children[c] == nil {
// 				curr.children[c] = newTrieNode(idx + 1)
// 			}
// 			curr = curr.children[c]
// 		}
// 	}
// 	curr = root
// 	var index = -1
// 	for _, c := range searchWord {
// 		if curr.children[c] == nil {
// 			index = -1
// 			break
// 		}
// 		curr = curr.children[c]
// 		index = curr.index
// 	}
// 	return index
// }
