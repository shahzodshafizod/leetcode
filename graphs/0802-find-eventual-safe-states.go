package graphs

// https://leetcode.com/problems/find-eventual-safe-states/

// Approach: Depth-First Search
// Time: O(v+e), v=# of nodes, e=# of edges
// Space: O(v)
func eventualSafeNodes(graph [][]int) []int {
	var n = len(graph)
	var safe = make([]*bool, n)
	var seen = make([]bool, n)
	var dfs func(int) bool
	dfs = func(node int) bool {
		if seen[node] {
			return false
		}
		if safe[node] != nil {
			return *safe[node]
		}
		seen[node] = true
		safe[node] = new(bool)
		*safe[node] = true
		for _, next := range graph[node] {
			*safe[node] = *safe[node] && dfs(next)
		}
		seen[node] = false
		return *safe[node]
	}
	var result = make([]int, 0)
	for node := 0; node < n; node++ {
		if dfs(node) {
			result = append(result, node)
		}
	}
	return result
}

