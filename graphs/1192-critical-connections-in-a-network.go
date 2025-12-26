package graphs

// https://leetcode.com/problems/critical-connections-in-a-network/

// Approach: Tarjan's Algorithm for finding bridges
// Use DFS with discovery time and low link values
// Time: O(V + E)
// Space: O(V + E)
func criticalConnections(n int, connections [][]int) [][]int {
	// Build adjacency list
	graph := make([][]int, n)

	for _, conn := range connections {
		u, v := conn[0], conn[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	dfn := make([]int, n)
	low := make([]int, n)

	for i := range dfn {
		dfn[i] = -1
	}

	bridges := [][]int{}
	timer := 0

	var dfs func(int, int)

	dfs = func(node, parent int) {
		dfn[node] = timer
		low[node] = timer
		timer++

		for _, neighbor := range graph[node] {
			if neighbor == parent {
				continue
			}

			if dfn[neighbor] == -1 {
				dfs(neighbor, node)
				low[node] = min(low[node], low[neighbor])

				// Bridge found
				if low[neighbor] > dfn[node] {
					bridges = append(bridges, []int{node, neighbor})
				}
			} else {
				low[node] = min(low[node], dfn[neighbor])
			}
		}
	}

	dfs(0, -1)

	return bridges
}
