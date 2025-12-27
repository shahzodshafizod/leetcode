package graphs

import "math"

// https://leetcode.com/problems/minimum-degree-of-a-connected-trio-in-a-graph/

// Approach #2: Optimized with Adjacency Matrix
// Time: O(m * n) where m = number of edges
// Space: O(n^2) - adjacency matrix
func minTrioDegree(n int, edges [][]int) int {
	// Build adjacency matrix for O(1) edge lookup
	adjMatrix := make([][]bool, n+1)
	for i := range adjMatrix {
		adjMatrix[i] = make([]bool, n+1)
	}

	degree := make([]int, n+1)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjMatrix[u][v] = true
		adjMatrix[v][u] = true
		degree[u]++
		degree[v]++
	}

	minDegree := math.MaxInt

	// For each edge (u, v), find common neighbors to form trios
	for _, edge := range edges {
		u, v := edge[0], edge[1]

		// Ensure u < v for consistent ordering
		if u > v {
			u, v = v, u
		}

		// Find all nodes w that are connected to both u and v
		for w := 1; w <= n; w++ {
			if w != u && w != v && adjMatrix[u][w] && adjMatrix[v][w] {
				// Found a trio (u, v, w)
				// Degree = sum of degrees - 6 (3 internal edges * 2)
				trioDegree := degree[u] + degree[v] + degree[w] - 6
				if trioDegree < minDegree {
					minDegree = trioDegree
				}
			}
		}
	}

	if minDegree == math.MaxInt {
		return -1
	}

	return minDegree
}

// // Approach #1: Brute-Force (Check all possible trios)
// // Time: O(n^3) - three nested loops to check all node combinations
// // Space: O(n^2) - adjacency matrix
// func minTrioDegree(n int, edges [][]int) int {
// 	// Build adjacency list and count degrees
// 	adj := make([]map[int]bool, n+1)
// 	for i := range adj {
// 		adj[i] = make(map[int]bool)
// 	}
// 	degree := make([]int, n+1)

// 	for _, edge := range edges {
// 		u, v := edge[0], edge[1]
// 		adj[u][v] = true
// 		adj[v][u] = true
// 		degree[u]++
// 		degree[v]++
// 	}

// 	minDegree := math.MaxInt

// 	// Check all possible trios (i, j, k) where i < j < k
// 	for i := 1; i <= n; i++ {
// 		for j := i + 1; j <= n; j++ {
// 			if !adj[i][j] {
// 				continue
// 			}
// 			for k := j + 1; k <= n; k++ {
// 				// Check if (i, j, k) forms a trio
// 				if adj[i][k] && adj[j][k] {
// 					// Calculate degree of this trio
// 					// degree = (deg(i) + deg(j) + deg(k)) - 6
// 					// The -6 removes the 3 internal edges (counted twice each)
// 					trioDegree := degree[i] + degree[j] + degree[k] - 6
// 					if trioDegree < minDegree {
// 						minDegree = trioDegree
// 					}
// 				}
// 			}
// 		}
// 	}

// 	if minDegree == math.MaxInt {
// 		return -1
// 	}
// 	return minDegree
// }
