package unionfinds

import "sort"

// https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/

// Approach: Union-Find + Sorting
// Time: O((E + Q) log(E + Q)) - sorting edges and queries
// Space: O(N + Q) - parent array and result array
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// Union-Find data structure
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int

	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // Path compression
		}

		return parent[x]
	}

	union := func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			parent[py] = px
		}
	}

	// Sort edges by distance (ascending)
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	// Create indexed queries: [p, q, limit, originalIndex]
	type Query struct {
		p, q, limit, idx int
	}

	indexedQueries := make([]Query, len(queries))
	for i, q := range queries {
		indexedQueries[i] = Query{q[0], q[1], q[2], i}
	}

	// Sort queries by limit (ascending)
	sort.Slice(indexedQueries, func(i, j int) bool {
		return indexedQueries[i].limit < indexedQueries[j].limit
	})

	result := make([]bool, len(queries))
	edgeIdx := 0

	// Process queries in order of increasing limit
	for _, query := range indexedQueries {
		// Add all edges with distance < limit to Union-Find
		for edgeIdx < len(edgeList) && edgeList[edgeIdx][2] < query.limit {
			u, v := edgeList[edgeIdx][0], edgeList[edgeIdx][1]
			union(u, v)

			edgeIdx++
		}

		// Check if p and q are connected
		result[query.idx] = find(query.p) == find(query.q)
	}

	return result
}
