package graphs

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-edges-in-shortest-paths/

func findAnswer(n int, edges [][]int) []bool {
	graph := make([][][3]int, n)
	var a, b, w int
	for i, edge := range edges {
		a, b, w = edge[0], edge[1], edge[2]
		graph[a] = append(graph[a], [3]int{b, w, i})
		graph[b] = append(graph[b], [3]int{a, w, i})
	}
	prev := make([][][2]int, n) // {{node, edgeID}}
	best := make([]int, n)
	for i := range n {
		best[i] = (1 << 32) - 1
	}
	best[0] = 0
	pq := pkg.NewHeap(
		[][2]int{{0, 0}}, // {node, cost}
		func(x, y [2]int) bool { return x[1] < y[1] },
	)
	var node, cost, nei, weight, ncost, eidx int
	for pq.Len() > 0 { // Dijkstra's algorithm -> Time: O(E log E)
		item := heap.Pop(pq).([2]int)
		node, cost = item[0], item[1]
		if cost > best[node] {
			continue
		}
		for _, item := range graph[node] {
			nei, weight, eidx = item[0], item[1], item[2]
			ncost = weight + best[node]
			if ncost < best[nei] {
				best[nei] = ncost
				prev[nei] = [][2]int{{node, eidx}}
				heap.Push(pq, [2]int{nei, ncost})
			} else if ncost == best[nei] {
				prev[nei] = append(prev[nei], [2]int{node, eidx})
			}
		}
	}
	answer := make([]bool, len(edges))
	seen := make([]bool, n)
	var prv int
	var track func(node int)
	track = func(node int) { // DFS -> Time: O(V + E)
		if seen[node] {
			return
		}
		seen[node] = true
		for _, item := range prev[node] {
			prv, eidx = item[0], item[1]
			answer[eidx] = true
			track(prv)
		}
	}
	track(n - 1)
	return answer
}
