package graphs

// https://leetcode.com/problems/is-graph-bipartite/

// Approach #3: Depth-First Search
// Time: O(v+e), v=# of vertices (nodes), e=# of edges
// Space: O(v+e)
func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)

	var dfs func(node int, color int) bool

	dfs = func(node int, color int) bool {
		colors[node] = color
		for _, neighbor := range graph[node] {
			if colors[neighbor] == -color {
				continue
			} else if colors[neighbor] == color || !dfs(neighbor, -color) {
				return false
			}
		}

		return true
	}
	for node := range n {
		if colors[node] == 0 && !dfs(node, 1) {
			return false
		}
	}

	return true
}

// // Approach #2: Breadth-First Search
// // Time: O(v+e), v=# of vertices (nodes), e=# of edges
// // Space: O(v+e)
// func isBipartite(graph [][]int) bool {
// 	var n = len(graph)
// 	var colors = make([]int, n)
// 	var curr, next, size int
// 	var check = func(node int) bool {
// 		var queue = list.New()
// 		queue.PushBack(node)
// 		colors[node] = 1
// 		for size = queue.Len(); size > 0; size = queue.Len() {
// 			for ; size > 0; size-- {
// 				curr = queue.Remove(queue.Front()).(int)
// 				for _, next = range graph[curr] {
// 					if colors[next] == colors[curr] {
// 						return false
// 					} else if colors[next] == 0 {
// 						colors[next] = -1 * colors[curr]
// 						queue.PushBack(next)
// 					}
// 				}
// 			}
// 		}
// 		return true
// 	}
// 	for node := 0; node < n; node++ {
// 		if colors[node] == 0 && !check(node) {
// 			return false
// 		}
// 	}
// 	return true
// }

// // Approach #1: Union Find
// // Time: O(v+e), v=# of vertices, e=# of edges
// // Space: O(v)
// func isBipartite(graph [][]int) bool {
// 	var n = len(graph)
// 	var parent = make([]int, n)
// 	for node := 0; node < n; node++ {
// 		parent[node] = node
// 	}
// 	var find func(x int) int
// 	find = func(x int) int {
// 		if parent[x] != x {
// 			parent[x] = find(parent[x])
// 		}
// 		return parent[x]
// 	}
// 	var union = func(x int, y int) {
// 		px := find(x)
// 		py := find(y)
// 		if px != py {
// 			parent[py] = px
// 		}
// 	}
// 	for node := 0; node < n; node++ {
// 		for _, neighbor := range graph[node] {
// 			if find(node) == find(neighbor) {
// 				return false
// 			}
// 			union(graph[node][0], neighbor)
// 		}
// 	}
// 	return true
// }
