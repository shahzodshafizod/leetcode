package design

// https://leetcode.com/problems/prefix-and-suffix-search/

// Time: O(N∗M)
// Space: O(N∗M)
type WordFilter struct {
	index    int
	children map[byte]*WordFilter
}

func NewWordFilter(words []string) WordFilter {
	root := WordFilter{children: make(map[byte]*WordFilter)}

	var wordLen, newWordLen int

	var c byte

	for index := range words {
		wordLen = len(words[index])
		word := words[index] + "#" + words[index]
		newWordLen = 2*wordLen + 1

		for i := 0; i < wordLen; i++ {
			curr := &root

			for j := i; j < newWordLen; j++ {
				c = word[j]
				if curr.children[c] == nil {
					curr.children[c] = &WordFilter{children: make(map[byte]*WordFilter)}
				}

				curr = curr.children[c]
				curr.index = index
			}
		}
	}

	return root
}

func (w *WordFilter) F(pref string, suff string) int {
	word := suff + "#" + pref
	curr := w

	var c byte
	for idx := range word {
		c = word[idx]
		if curr.children[c] == nil {
			return -1
		}

		curr = curr.children[c]
	}

	return curr.index
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
