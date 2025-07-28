package hashes

import "strings"

// https://leetcode.com/problems/unique-email-addresses/

// Approach #2: Hash Set
// Time: O(n*m), n=len(emails), m=len(emails[i])
// Space: O(n)
func numUniqueEmails(emails []string) int {
	receivers := make(map[string]struct{})

	for _, email := range emails {
		parts := strings.Split(email, "@")
		parts[0] = strings.Split(parts[0], "+")[0]
		parts[0] = strings.ReplaceAll(parts[0], ".", "")
		receivers[strings.Join(parts, "@")] = struct{}{}
	}

	return len(receivers)
}

// // Approach #1: Trie
// // Time: O(n*m), n=len(emails), m=len(emails[i])
// // Space: O(n*m)
// func numUniqueEmails(emails []string) int {
// 	type TrieNode struct {
// 		children map[rune]*TrieNode
// 		end      bool
// 	}
// 	var NewTrieNode = func() *TrieNode {
// 		return &TrieNode{children: make(map[rune]*TrieNode)}
// 	}
// 	const (
// 		LOCAL = iota
// 		SKIP
// 		DOMAIN
// 	)
// 	var count = 0
// 	var root = NewTrieNode()
// 	for _, email := range emails {
// 		var curr = root
// 		var part = LOCAL
// 		for _, c := range email {
// 			if part != DOMAIN {
// 				switch c {
// 				case '.':
// 					continue
// 				case '+':
// 					part = SKIP
// 				case '@':
// 					part = DOMAIN
// 				}
// 			}
// 			if part == SKIP {
// 				continue
// 			}
// 			if curr.children[c] == nil {
// 				curr.children[c] = NewTrieNode()
// 			}
// 			curr = curr.children[c]
// 		}
// 		if !curr.end {
// 			count++
// 		}
// 		curr.end = true
// 	}
// 	return count
// }
