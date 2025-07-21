package tries

// https://leetcode.com/problems/longest-common-suffix-queries/

// M=len(wordsContainer)
// C=len(wordsContainer[i])
// N=len(wordsQuery)
// Q=len(wordsQuery[i])
// Time: O(M*C + N*Q)
// Space: O(M*C)
func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	type Trie struct {
		children [26]*Trie
		index    int
		length   int
	}
	newTrie := func() *Trie {
		return &Trie{
			index:  -1,
			length: 5001,
		}
	}
	buildTrie := func() *Trie {
		root := newTrie()
		for idx, word := range wordsContainer {
			curr := root
			length := len(word)
			if length < curr.length {
				curr.length, curr.index = length, idx
			}
			for c := len(word) - 1; c >= 0; c-- {
				if curr.children[word[c]-'a'] == nil {
					curr.children[word[c]-'a'] = newTrie()
				}
				curr = curr.children[word[c]-'a']
				if length < curr.length {
					curr.length, curr.index = length, idx
				}
			}
		}
		return root
	}
	root := buildTrie()
	ans := make([]int, len(wordsQuery))
	for idx, word := range wordsQuery {
		curr := root
		ans[idx] = curr.index
		for c := len(word) - 1; c >= 0; c-- {
			if curr.children[word[c]-'a'] == nil {
				break
			}
			curr = curr.children[word[c]-'a']
			ans[idx] = curr.index
		}
	}
	return ans
}
