package graphs

import "sort"

// https://leetcode.com/problems/maximum-score-of-a-node-sequence/

// Approach: Graph + Greedy (Store top 3 neighbors for each node)
// Time: O(E * 9) = O(E) - for each edge, check up to 9 combinations
// Space: O(n * 3) = O(n) - store top 3 neighbors per node
func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)

	// Build adjacency list
	neighbors := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		neighbors[u] = append(neighbors[u], v)
		neighbors[v] = append(neighbors[v], u)
	}

	// Keep only top 3 neighbors for each node (sorted by score descending)
	topNeighbors := make([][]int, n)
	for node := range n {
		// Sort neighbors by their scores in descending order
		sort.Slice(neighbors[node], func(i, j int) bool {
			return scores[neighbors[node][i]] > scores[neighbors[node][j]]
		})

		// Keep only top 3
		if len(neighbors[node]) > 3 {
			topNeighbors[node] = neighbors[node][:3]
		} else {
			topNeighbors[node] = neighbors[node]
		}
	}

	maxScore := -1

	// For each edge (b, c), try to form sequence: a - b - c - d
	for _, edge := range edges {
		b, c := edge[0], edge[1]

		// Try all combinations of top neighbors
		for _, a := range topNeighbors[b] {
			// a must not be c
			if a == c {
				continue
			}

			for _, d := range topNeighbors[c] {
				// d must not be b or a
				if d == b || d == a {
					continue
				}

				// Valid sequence: a - b - c - d
				currentScore := scores[a] + scores[b] + scores[c] + scores[d]
				if currentScore > maxScore {
					maxScore = currentScore
				}
			}
		}
	}

	return maxScore
}
