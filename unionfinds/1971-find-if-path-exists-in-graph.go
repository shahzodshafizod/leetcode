package unionfinds

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/find-if-path-exists-in-graph/

func validPath(n int, edges [][]int, source int, destination int) bool {
	uf := pkg.NewUnionFind(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}

	return uf.Find(source) == uf.Find(destination)
}

// func validPath(n int, edges [][]int, source int, destination int) bool {
// 	var adjList = make(map[int][]int)
// 	for _, edge := range edges {
// 		u, v := edge[0], edge[1]
// 		adjList[u] = append(adjList[u], v)
// 		adjList[v] = append(adjList[v], u)
// 	}
// 	var seen = make(map[int]bool)
// 	var dfs func(curr int) bool
// 	dfs = func(curr int) bool {
// 		if curr == destination {
// 			return true
// 		}
// 		seen[source] = true
// 		for _, next := range adjList[curr] {
// 			if !seen[next] && dfs(next) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	return dfs(source)
// }
