package tries

import "strings"

// https://leetcode.com/problems/replace-words/

type trie struct {
	end      bool
	children map[rune]*trie
}

func (t *trie) insert(word *string) {
	var current = t
	for _, c := range *word {
		child := current.children[c]
		if child == nil {
			child = &trie{children: make(map[rune]*trie)}
			current.children[c] = child
		}
		current = child
	}
	current.end = true
}

func (t *trie) search(word *string) int {
	var current = t
	for idx, c := range *word {
		if current.end {
			return idx
		}
		current = current.children[c]
		if current == nil {
			break
		}
	}
	return -1
}

// d: len(dictionary)
// s: len(words) - number of words in the sentence
// w: max(len(words[i]))
// time: O(d*w + s*w)
// space: O(d*w + s*w*26) = O(d*w + s*w)
// each node can store up to 26 pointers
func replaceWords(dictionary []string, sentence string) string {
	var trie = &trie{children: make(map[rune]*trie)}
	for idx := range dictionary {
		trie.insert(&dictionary[idx])
	}
	var words = strings.Split(sentence, " ")
	for idx := range words {
		rootEnd := trie.search(&words[idx])
		if rootEnd > 0 {
			words[idx] = words[idx][:rootEnd]
		}
	}
	return strings.Join(words, " ")
}

// // d: len(dictionary)
// // s: len(words) - number of words in the sentence
// // w: max(len(words[i]))
// // time: O(d*w + s*w^2)
// // space: O(d*w + s*w)
// func replaceWords(dictionary []string, sentence string) string {
// 	var roots = make(map[string]bool)
// 	for _, root := range dictionary {
// 		roots[root] = true
// 	}
// 	var words = strings.Split(sentence, " ")
// 	for idx, word := range words {
// 		for end := range word {
// 			if roots[word[:end]] {
// 				words[idx] = word[:end]
// 				break
// 			}
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
