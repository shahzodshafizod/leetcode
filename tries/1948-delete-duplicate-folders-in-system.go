package tries

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode.com/problems/delete-duplicate-folders-in-system/

func deleteDuplicateFolder(paths [][]string) [][]string {
	type node map[string]node
	root := node{}
	for _, path := range paths {
		curr := root
		for _, folder := range path {
			if curr[folder] == nil {
				curr[folder] = node{}
			}
			curr = curr[folder]
		}
	}
	seen := make(map[string][]node)
	var serialize func(curr node) string
	serialize = func(curr node) string {
		if len(curr) == 0 {
			return ""
		}
		var ss []string
		for folder, next := range curr {
			ss = append(ss, fmt.Sprintf("%s(%s)", folder, serialize(next)))
		}
		sort.Strings(ss)
		key := strings.Join(ss, "")
		seen[key] = append(seen[key], curr)
		return key
	}
	serialize(root)
	for _, nodes := range seen {
		if len(nodes) > 1 {
			for _, x := range nodes {
				x["$"] = node{}
			}
		}
	}
	var cleaned [][]string
	var dfs func(curr node, path []string)
	dfs = func(curr node, path []string) {
		for folder, next := range curr {
			if next["$"] == nil {
				n := len(path) + 1
				newPath := make([]string, n)
				copy(newPath, path)
				newPath[n-1] = folder
				cleaned = append(cleaned, newPath)
				dfs(next, newPath)
			}
		}
	}
	dfs(root, []string{})
	return cleaned
}
