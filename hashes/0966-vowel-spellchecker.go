package hashes

import "strings"

// https://leetcode.com/problems/vowel-spellchecker/

// Approach #2: Hash Table
// Time: O(n)
// Space: O(n)
func spellchecker(wordlist []string, queries []string) []string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	devowel := func(word string) string {
		var sb strings.Builder
		sb.Grow(len(word))

		var err error

		for idx := range word {
			if vowels[word[idx]] {
				err = sb.WriteByte('*')
			} else {
				err = sb.WriteByte(word[idx])
			}

			_ = err
		}

		return sb.String()
	}
	orig := make(map[string]struct{})
	lower := make(map[string]string)
	vowel := make(map[string]string)

	for i := len(wordlist) - 1; i >= 0; i-- {
		orig[wordlist[i]] = struct{}{}
		lower[strings.ToLower(wordlist[i])] = wordlist[i]
		vowel[devowel(strings.ToLower(wordlist[i]))] = wordlist[i]
	}

	answer := make([]string, 0, len(queries))

	var (
		val   string
		found bool
	)
	for _, word := range queries {
		if _, found = orig[word]; found {
			answer = append(answer, word)
		} else if val, found = lower[strings.ToLower(word)]; found {
			answer = append(answer, val)
		} else if val, found = vowel[devowel(strings.ToLower(word))]; found {
			answer = append(answer, val)
		} else {
			answer = append(answer, "")
		}
	}

	return answer
}

// // Approach #1: Tries
// // Time: O(n*m)
// // Space: O(n*m)
// func spellchecker(wordlist []string, queries []string) []string {
// 	type node struct {
// 		word     string
// 		children map[rune]*node
// 	}

// 	newNode := func() *node {
// 		return &node{children: make(map[rune]*node)}
// 	}
// 	insert := func(curr *node, word string, orig string) {
// 		for _, c := range word {
// 			if curr.children[c] == nil {
// 				curr.children[c] = newNode()
// 			}

// 			curr = curr.children[c]
// 		}

// 		curr.word = orig
// 	}
// 	find := func(curr *node, word string) string {
// 		for _, c := range word {
// 			if curr.children[c] == nil {
// 				return ""
// 			}

// 			curr = curr.children[c]
// 		}

// 		return curr.word
// 	}
// 	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
// 	devowel := func(word string) string {
// 		bword := []byte(strings.ToLower(word))
// 		for i, c := range bword {
// 			if vowels[c] {
// 				bword[i] = '*'
// 			}
// 		}

// 		return string(bword)
// 	}

// 	orig, lower, vowel := newNode(), newNode(), newNode()
// 	for i := len(wordlist) - 1; i >= 0; i-- {
// 		insert(orig, wordlist[i], wordlist[i])
// 		insert(lower, strings.ToLower(wordlist[i]), wordlist[i])
// 		insert(vowel, devowel(wordlist[i]), wordlist[i])
// 	}

// 	type tuple struct {
// 		root *node
// 		word string
// 	}

// 	answer := make([]string, len(queries))

// 	var ans string

// 	for i, word := range queries {
// 		for _, tuple := range []*tuple{{orig, word}, {lower, strings.ToLower(word)}, {vowel, devowel(word)}} {
// 			ans = find(tuple.root, tuple.word)
// 			if len(ans) != 0 {
// 				answer[i] = ans

// 				break
// 			}
// 		}
// 	}

// 	return answer
// }
