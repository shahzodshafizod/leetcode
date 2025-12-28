package trees

// https://leetcode.com/problems/frog-position-after-t-seconds/

// Approach: Optimized - DFS with early termination
// Use DFS to find target path, calculate probability along the way
// Key insight: probability = product of (1/choices) at each step
// Can only be at target if: (1) exactly at time t, or (2) stuck at target before t
// N = number of nodes, E = number of edges
// Time: O(N + E) - worst case visit all nodes, but often terminates early
// Space: O(N) - recursion stack depth (tree height)
func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// Build adjacency list
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var dfs func(node, parent, time int, prob float64) float64

	dfs = func(node, parent, time int, prob float64) float64 {
		// Count unvisited neighbors (exclude parent)
		numNeighbors := 0

		for _, neighbor := range graph[node] {
			if neighbor != parent {
				numNeighbors++
			}
		}

		// If we found the target
		if node == target {
			// Case 1: We're exactly at time t
			if time == t {
				return prob
			}
			// Case 2: We reached target before t, but can't move (stuck here)
			if numNeighbors == 0 {
				return prob
			}
			// Case 3: We reached target but will have to move away
			return 0.0
		}

		// If time is up or no neighbors to explore, can't reach target
		if time >= t || numNeighbors == 0 {
			return 0.0
		}

		// Explore neighbors with equal probability distribution
		for _, neighbor := range graph[node] {
			if neighbor != parent {
				result := dfs(neighbor, node, time+1, prob/float64(numNeighbors))
				if result > 0 {
					return result
				}
			}
		}

		return 0.0
	}

	return dfs(1, -1, 0, 1.0)
}
