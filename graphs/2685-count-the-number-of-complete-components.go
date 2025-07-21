package graphs

// https://leetcode.com/problems/count-the-number-of-complete-components/

// Approach #2: Depth-First Search
// Time: O(E+2V) = O(E+V)
// Space: O(E+3V) = O(E+V)
func countCompleteComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	var a, b int
	for idx := range edges {
		a, b = edges[idx][0], edges[idx][1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	seen := make([]bool, n)
	var dfs func(node int) []int
	dfs = func(node int) []int {
		if seen[node] {
			return []int{}
		}
		seen[node] = true
		vertices := []int{node}
		for _, nei := range graph[node] {
			vertices = append(vertices, dfs(nei)...)
		}
		return vertices
	}
	var complete int
	count := 0
	for node := 0; node < n; node++ {
		if !seen[node] {
			vertices := dfs(node)
			complete = 1
			for _, v := range vertices {
				if len(graph[v]) != len(vertices)-1 {
					complete = 0
					break
				}
			}
			count += complete
		}
	}
	return count
}

// // Approach #1: Depth-First Search + Union Find
// // Time: O(2E+2V) = O(E+V)
// // Space: O(E+4V) = O(E+V)
// func countCompleteComponents(n int, edges [][]int) int {
// 	var uf = pkg.NewUnionFindRanked(n)
// 	var graph = make([][]int, n)
// 	var a, b int
// 	for idx := range edges {
// 		a, b = edges[idx][0], edges[idx][1]
// 		uf.Union(a, b)
// 		graph[a] = append(graph[a], b)
// 		graph[b] = append(graph[b], a)
// 	}
// 	var compEdges = make(map[int]int)
// 	var root int
// 	for idx := range edges {
// 		root = uf.Find(edges[idx][0])
// 		compEdges[root]++
// 	}
// 	var seen = make([]bool, n)
// 	var dfs func(node int) int
// 	dfs = func(node int) int {
// 		if seen[node] {
// 			return 0
// 		}
// 		seen[node] = true
// 		var count = 1
// 		for _, nei := range graph[node] {
// 			count += dfs(nei)
// 		}
// 		return count
// 	}
// 	var count = 0
// 	var compNodes int
// 	for node := 0; node < n; node++ {
// 		if !seen[node] {
// 			compNodes = dfs(node)
// 			if compNodes*(compNodes-1)/2 == compEdges[uf.Find(node)] {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
