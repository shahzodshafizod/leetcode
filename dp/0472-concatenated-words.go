package dp

import "sort"

// https://leetcode.com/problems/concatenated-words/

// Approach #3: Bottom-Up Dynamic Programming + Hash Set
// Time: O(NLogN+N∗M^2), N=len(words), M=max(len(words[i]))
// Space: O(N+M)
func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	set := make(map[string]bool)
	concatenated := make([]string, 0)
	for _, word := range words {
		n := len(word)
		dp := make([]bool, n+1)
		dp[0] = true
		for len := 1; len <= n; len++ {
			for j := len - 1; j >= 0; j-- {
				if dp[j] && set[word[j:len]] {
					dp[len] = true
					break
				}
			}
		}
		if dp[n] {
			concatenated = append(concatenated, word)
		}
		set[word] = true
	}
	return concatenated
}

// // Approach #2: Depth-First Search + Hash Set
// // Time: O(n*www), n=len(words), w=max(len(words[i]))
// // Space: O(n)
// func findAllConcatenatedWordsInADict(words []string) []string {
// 	var set = make(map[string]bool)
// 	for _, word := range words {
// 		set[word] = true
// 	}
// 	var memo = make(map[string]*bool)
// 	var dfs func(word string) bool
// 	dfs = func(word string) bool {
// 		if memo[word] != nil {
// 			return *memo[word]
// 		}
// 		memo[word] = new(bool)
// 		for idx, n := 1, len(word); idx < n; idx++ {
// 			prefix := word[:idx]
// 			suffix := word[idx:]
// 			if set[prefix] && (set[suffix] || dfs(suffix)) {
// 				*memo[word] = true
// 				break
// 			}
// 		}
// 		return *memo[word]
// 	}
// 	var concatenated = make([]string, 0)
// 	for _, word := range words {
// 		if dfs(word) {
// 			concatenated = append(concatenated, word)
// 		}
// 	}
// 	return concatenated
// }

// // Approach #1: Depth-First Search + Trie
// // Time: O(n*www), n=len(words), w=max(len(words[i]))
// // Space: O(n*w)
// func findAllConcatenatedWordsInADict(words []string) []string {
// 	type TrieNode struct {
// 		end      bool
// 		children map[byte]*TrieNode
// 	}
// 	var newTrieNode = func() *TrieNode {
// 		return &TrieNode{children: make(map[byte]*TrieNode)}
// 	}
// 	var root = newTrieNode()
// 	var insert = func(word string) {
// 		var curr *TrieNode
// 		var c byte
// 		curr = root
// 		for idx := range word {
// 			c = word[idx]
// 			if curr.children[c] == nil {
// 				curr.children[c] = newTrieNode()
// 			}
// 			curr = curr.children[c]
// 		}
// 		curr.end = true
// 	}
// 	var maxLen = 0
// 	for _, word := range words {
// 		maxLen = max(maxLen, len(word))
// 	}
// 	var memo = make([][]*bool, len(words))
// 	for idx := range memo {
// 		memo[idx] = make([]*bool, maxLen)
// 	}
// 	var dfs func(row int, col int) bool
// 	dfs = func(row int, col int) bool {
// 		var n = len(words[row])
// 		if col == n {
// 			return true
// 		}
// 		if memo[row][col] != nil {
// 			return *memo[row][col]
// 		}
// 		memo[row][col] = new(bool)
// 		var curr = root
// 		var result = false
// 		for idx := col; idx < n; idx++ {
// 			if curr.children[words[row][idx]] == nil {
// 				break
// 			}
// 			curr = curr.children[words[row][idx]]
// 			if curr.end && dfs(row, idx+1) {
// 				result = true
// 				break
// 			}
// 		}
// 		*memo[row][col] = result
// 		return result
// 	}
// 	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
// 	var concatenated = make([]string, 0)
// 	for row, word := range words {
// 		if dfs(row, 0) {
// 			concatenated = append(concatenated, word)
// 		}
// 		insert(word)
// 	}
// 	return concatenated
// }
