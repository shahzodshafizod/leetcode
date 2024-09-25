package tries

// https://leetcode.com/problems/sum-of-prefix-scores-of-strings/

// Approach: Trie
// N=len(words)
// L=len(words[i])
// Time: O(N*L)
// Space: O(N*L)
func sumPrefixScores(words []string) []int {
	type TrieNode struct {
		children map[rune]*TrieNode
		count    int
	}
	var newTrieNode = func() *TrieNode {
		return &TrieNode{children: make(map[rune]*TrieNode)}
	}
	var root = newTrieNode()
	for _, word := range words {
		curr := root
		for _, c := range word {
			if curr.children[c] == nil {
				curr.children[c] = newTrieNode()
			}
			curr = curr.children[c]
			curr.count++
		}
	}
	var answer = make([]int, len(words))
	for idx, word := range words {
		curr := root
		for _, c := range word {
			curr = curr.children[c]
			answer[idx] += curr.count
		}
	}
	return answer
}

// // Approach: Hashmap+Counting
// // N=len(words)
// // L=len(words[i])
// // Time: O(N*L^2)
// // Space: O(N*L)
// func sumPrefixScores(words []string) []int {
// 	var prefCount = make(map[string]int)
// 	for _, word := range words { // O(N)
// 		for idx := range word { // O(L)
// 			prefCount[word[:idx+1]]++ // O(L)
// 		}
// 	}
// 	var answer = make([]int, len(words))
// 	for i, word := range words {
// 		for idx := range word {
// 			answer[i] += prefCount[word[:idx+1]]
// 		}
// 	}
// 	return answer
// }
