package tries

// https://leetcode.com/problems/longest-common-prefix/

// Approach: Brute-Force
// N = len(strs)
// Time: O(N x len(prefix))
// Space: O(1)
func longestCommonPrefix(strs []string) string {
	idx := 0
	exit := false
	for !exit {
		for _, word := range strs {
			if idx == len(word) || word[idx] != strs[0][idx] {
				exit = true
				idx--
				break
			}
		}
		idx++
	}
	return strs[0][:idx]
}

// // Approach: Trie
// func longestCommonPrefix(strs []string) string {
// 	type node struct {
// 		words    int
// 		children [26]*node
// 	}
// 	// insert
// 	var root = &node{}
// 	for _, word := range strs {
// 		var curr = root
// 		curr.words++
// 		for _, c := range word {
// 			if curr.children[c-'a'] == nil {
// 				curr.children[c-'a'] = &node{}
// 			}
// 			curr = curr.children[c-'a']
// 			curr.words++
// 		}
// 	}
// 	var n = len(strs)
// 	var search func(curr *node, prefix string) string
// 	search = func(curr *node, prefix string) string {
// 		if curr.words != n {
// 			return ""
// 		}
// 		for c, child := range curr.children {
// 			if child == nil {
// 				continue
// 			}
// 			pref := search(child, prefix+string(rune('a'+c)))
// 			if len(pref) > len(prefix) {
// 				prefix = pref
// 			}
// 		}
// 		return prefix
// 	}
// 	return search(root, "")
// }
