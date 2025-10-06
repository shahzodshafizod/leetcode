package graphs

// https://leetcode.com/problems/similar-string-groups/

// Approach #2: Depth-First Search
// Time: O(nnm), n=len(strs), m=len(strs[i])
// Space: O(nm)
func numSimilarGroups(strs []string) int {
	n := len(strs)
	similar := func(i int, j int) bool {
		var count int

		for k := len(strs[i]) - 1; k >= 0 && count <= 2; k-- {
			if strs[i][k] != strs[j][k] {
				count++
			}
		}

		return count <= 2
	}
	seen := make([]bool, n)

	var dfs func(i int)

	dfs = func(i int) {
		seen[i] = true
		for j := range n {
			if !seen[j] && similar(i, j) {
				dfs(j)
			}
		}
	}

	var groups int

	for i := range n {
		if !seen[i] {
			groups++

			dfs(i)
		}
	}

	return groups
}

// // Approach #1: Union Find
// // Time: O(nnm), n=len(strs), m=len(strs[i])
// // Space: O(nm)
// func numSimilarGroups(strs []string) int {
// 	n := len(strs)

// 	root := make([]int, n)
// 	for i := range n {
// 		root[i] = i
// 	}

// 	groups := n

// 	var find func(x int) int

// 	find = func(x int) int {
// 		if root[x] != x {
// 			root[x] = find(root[x])
// 		}

// 		return root[x]
// 	}
// 	union := func(x int, y int) {
// 		rx, ry := find(x), find(y)
// 		if rx < ry {
// 			root[ry] = rx
// 			groups--
// 		} else if rx > ry {
// 			root[rx] = ry
// 			groups--
// 		}
// 	}

// 	similar := func(i int, j int) bool {
// 		var count int

// 		for k := len(strs[i]) - 1; k >= 0; k-- {
// 			if strs[i][k] != strs[j][k] {
// 				count++
// 			}
// 		}

// 		return count == 2 || count == 0
// 	}
// 	for i := range n {
// 		for j := i + 1; j < n; j++ {
// 			if similar(i, j) {
// 				union(i, j)
// 			}
// 		}
// 	}

// 	return groups
// }
