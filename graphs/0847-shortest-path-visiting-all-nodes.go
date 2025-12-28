package graphs

// https://leetcode.com/problems/shortest-path-visiting-all-nodes/

// Approach: BFS with Bitmask to track visited nodes
// Use state = (current_node, visited_mask) to represent current position
// visited_mask is a bitmask where bit i is 1 if node i has been visited
// Start BFS from all nodes simultaneously since we can start anywhere
// N = number of nodes (1 <= N <= 12, so bitmask fits in int)
// Time: O(N * 2^N) - at most N * 2^N states (node × visited combinations)
// Space: O(N * 2^N) - queue and visited set
func shortestPathLength(graph [][]int) int {
	n := len(graph)

	// Special case: single node
	if n == 1 {
		return 0
	}

	// Target: all nodes visited (all bits set)
	target := (1 << n) - 1

	// State: (node, mask, distance)
	type State struct {
		node int
		mask int
		dist int
	}

	queue := []State{}
	visited := make(map[[2]int]bool)

	// Start from all nodes (we can choose any starting point)
	for i := range n {
		mask := 1 << i // Only node i visited
		queue = append(queue, State{i, mask, 0})
		visited[[2]int{i, mask}] = true
	}

	// BFS to find shortest path
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		node, mask, dist := state.node, state.mask, state.dist

		// Check if all nodes visited
		if mask == target {
			return dist
		}

		// Explore neighbors
		for _, neighbor := range graph[node] {
			newMask := mask | (1 << neighbor) // Mark neighbor as visited
			key := [2]int{neighbor, newMask}

			if !visited[key] {
				visited[key] = true

				queue = append(queue, State{neighbor, newMask, dist + 1})
			}
		}
	}

	return -1 // Should never reach here for connected graph
}
