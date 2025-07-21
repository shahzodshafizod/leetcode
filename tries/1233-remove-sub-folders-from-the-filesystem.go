package tries

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/

// Approach #2: Tries
// Time: O(nl), n=len(folder), l=len(folder[i])
// Space: O(nl)
func removeSubfolders(folder []string) []string {
	type node map[string]node
	sort.Strings(folder)
	root := make(node)
	var curr node
	var res []string
	for _, f := range folder {
		curr = root
		for _, word := range strings.Split(f, "/") {
			if curr["$"] != nil {
				break
			}
			if curr[word] == nil {
				curr[word] = make(node)
			}
			curr = curr[word]
		}
		if curr["$"] == nil {
			res = append(res, f)
			curr["$"] = make(node)
		}
	}
	return res
}

// // Approach #1: Brute-Force (Stack)
// // Time: O(nll), n=len(folder), l=len(folder[i])
// // Space: O(n)
// func removeSubfolders(folder []string) []string {
// 	sort.Strings(folder)
// 	var stack []string
// 	var size int
// 	for _, f := range folder {
// 		if size == 0 || !strings.HasPrefix(f, stack[size-1]+"/") {
// 			stack = append(stack, f)
// 			size++
// 		}
// 	}
// 	return stack
// }
