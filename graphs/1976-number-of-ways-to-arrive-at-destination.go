package graphs

import (
	"container/heap"
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/

// Approach: Dijkstra's Algorithm
// Time: O((V+E)logV)
// Space: O(V+E)
func countPaths1976(n int, roads [][]int) int {
	dist := make([]int, n)
	for idx := 1; idx < n; idx++ {
		dist[idx] = math.MaxInt
	}
	count := make([]int, n)
	count[0] = 1
	graph := make([][][2]int, n)
	var u, v, t int
	for _, road := range roads {
		u, v, t = road[0], road[1], road[2]
		graph[u] = append(graph[u], [2]int{v, t})
		graph[v] = append(graph[v], [2]int{u, t})
	}
	queue := pkg.NewHeap(
		[][2]int{{0, 0}}, // node, time
		func(x, y [2]int) bool { return x[1] < y[1] },
	)
	var item [2]int
	var node, time, nn, nt int
	const MOD = int(1e9) + 7
	for queue.Len() > 0 {
		item = heap.Pop(queue).([2]int)
		node, time = item[0], item[1]
		if time > dist[n-1] {
			continue
		}
		for _, nei := range graph[node] {
			nn, nt = nei[0], nei[1]
			if time+nt < dist[nn] {
				dist[nn] = time + nt
				count[nn] = count[node]
				heap.Push(queue, [2]int{nn, dist[nn]})
			} else if time+nt == dist[nn] {
				count[nn] = (count[nn] + count[node]) % MOD
			}
		}
	}
	return count[n-1]
}
