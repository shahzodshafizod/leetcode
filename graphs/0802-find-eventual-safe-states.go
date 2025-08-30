package graphs

// https://leetcode.com/problems/find-eventual-safe-states/

// Approach: Depth-First Search
// Time: O(v+e), v=# of vertices (nodes), e=# of edges
// Space: O(v)
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	safe := make([]*bool, n)

	var dfs func(int) bool

	dfs = func(node int) bool {
		if safe[node] != nil {
			return *safe[node]
		}

		safe[node] = new(bool)

		for _, next := range graph[node] {
			if !dfs(next) {
				return false
			}
		}

		*safe[node] = true

		return true
	}
	result := make([]int, 0)

	for node := range n {
		if dfs(node) {
			result = append(result, node)
		}
	}

	return result
}
