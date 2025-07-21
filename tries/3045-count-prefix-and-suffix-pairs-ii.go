package tries

// https://leetcode.com/problems/count-prefix-and-suffix-pairs-ii/

func countPrefixSuffixPairs3045(words []string) int64 {
	type TrieNode struct {
		count    int
		children map[[2]byte]*TrieNode
	}
	NewTrieNode := func() *TrieNode {
		return &TrieNode{children: make(map[[2]byte]*TrieNode)}
	}
	root := NewTrieNode()
	var count int64 = 0
	var left, right int
	for wid := len(words) - 1; wid >= 0; wid-- {
		word := words[wid]
		left, right = 0, len(word)-1
		curr := root
		for right >= 0 {
			key := [2]byte{word[left], word[right]}
			if curr.children[key] == nil {
				curr.children[key] = NewTrieNode()
			}
			curr = curr.children[key]
			curr.count++
			left++
			right--
		}
		count += int64(curr.count - 1)
	}
	return count
}
