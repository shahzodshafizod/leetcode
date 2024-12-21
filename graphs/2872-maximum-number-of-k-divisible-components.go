package graphs

// https://leetcode.com/problems/maximum-number-of-k-divisible-components/

// Approach: Depth-First Search (Graph)
// Time: Time: O(n+e), n=# of nodes, e=# of edges
// Space: O(n+e)
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	var adj = make([][]int, n)
	var a, b int
	for idx := range edges {
		a, b = edges[idx][0], edges[idx][1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	var count = 0
	var dfs func(parent int, node int) int
	dfs = func(parent int, node int) int {
		var total = values[node]
		for _, next := range adj[node] {
			if next == parent {
				continue
			}
			total += dfs(node, next)
		}
		if total%k == 0 {
			count++
			total = 0
		}
		return total
	}
	dfs(-1, 0)
	return count
}

// // Approach: Breadth-First Search (Graph)
// // Time: Time: O(n+e), n=# of nodes, e=# of edges
// // Space: O(n+e)
// func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
// 	if n < 2 {
// 		return 1
// 	}
// 	var adj = make([][]int, n)
// 	var indegree = make([]int, n)
// 	var node1, node2 int
// 	for idx := range edges {
// 		node1, node2 = edges[idx][0], edges[idx][1]
// 		adj[node1] = append(adj[node1], node2)
// 		adj[node2] = append(adj[node2], node1)
// 		indegree[node1]++
// 		indegree[node2]++
// 	}
// 	var queue = make([]int, n)
// 	var idx, size = 0, 0
// 	for node := 0; node < n; node++ {
// 		if indegree[node] == 1 {
// 			queue[size] = node
// 			size++
// 		}
// 	}
// 	var count = 0
// 	var node, total int
// 	for idx < size {
// 		node = queue[idx]
// 		idx++
// 		indegree[node]--
// 		total = 0
// 		if values[node]%k == 0 {
// 			count++
// 		} else {
// 			total = values[node]
// 		}
// 		for _, next := range adj[node] {
// 			if indegree[next] == 0 {
// 				continue
// 			}
// 			indegree[next]--
// 			values[next] += total
// 			if indegree[next] == 1 {
// 				queue[size] = next
// 				size++
// 			}
// 		}
// 	}
// 	return count
// }
