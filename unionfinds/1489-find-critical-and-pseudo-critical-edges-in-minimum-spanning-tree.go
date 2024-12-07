package unionfinds

import (
	"sort"
)

// https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	var parent = make([]int, n)
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union = func(x int, y int) bool {
		var px, py = find(x), find(y)
		if px != py {
			parent[max(px, py)] = min(px, py)
		}
		return px != py
	}
	for idx := range edges {
		edges[idx] = append(edges[idx], idx)
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })
	var kruskal = func(skip int, pick int) int {
		for idx := 0; idx < n; idx++ {
			parent[idx] = idx
		}
		var weight = 0
		var count = n
		var node1, node2, cost int
		if pick >= 0 {
			node1, node2, cost = edges[pick][0], edges[pick][1], edges[pick][2]
			if union(node1, node2) {
				weight += cost
				count -= 1
			}
		}
		for idx := range edges {
			if idx == skip {
				continue
			}
			node1, node2, cost = edges[idx][0], edges[idx][1], edges[idx][2]
			if union(node1, node2) {
				weight += cost
				count -= 1
			}
		}
		if count != 1 {
			return 1_000_000
		}
		return weight
	}
	var critical, pseudoc = make([]int, 0), make([]int, 0)
	var mstWeight = kruskal(-1, -1)
	for idx := range edges {
		if kruskal(idx, -1) > mstWeight {
			critical = append(critical, edges[idx][3])
		} else if kruskal(-1, idx) == mstWeight {
			pseudoc = append(pseudoc, edges[idx][3])
		}
	}
	return [][]int{critical, pseudoc}
}
