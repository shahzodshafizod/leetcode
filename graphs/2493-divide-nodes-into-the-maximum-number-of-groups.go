package graphs

import "container/list"

// https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/

// Approach: Union-Find + Breadth-First Search
// Time: O(v*(v+e)), v=# of vertices, e=# of edges
// Space: O(v+e)
func magnificentSets(n int, edges [][]int) int {
	var parent = make([]int, n)
	var depth = make([]int, n)
	for node := 0; node < n; node++ {
		parent[node] = node
	}
	var find func(node int) int
	find = func(node int) int {
		if parent[node] != node {
			parent[node] = find(parent[node])
		}
		return parent[node]
	}
	var union = func(x int, y int) {
		px, py := find(x), find(y)
		if px == py {
			return
		}
		if depth[px] < depth[py] {
			px, py = py, px
		}
		parent[py] = px
		if depth[px] == depth[py] {
			depth[px]++
		}
	}
	var graph = make([][]int, n)
	var a, b int
	for idx := range edges {
		a, b = edges[idx][0]-1, edges[idx][1]-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		union(a, b) // groupping connected components
	}
	var curr, next int
	var countGroups = func(src int) int {
		var queue = list.New()
		queue.PushBack(src)
		var layers = make([]int, n)
		layers[src] = 1
		var layer = 0
		for size := queue.Len(); size > 0; size = queue.Len() {
			layer++
			for ; size > 0; size-- {
				curr = queue.Remove(queue.Front()).(int)
				for _, next = range graph[curr] {
					if layers[next] == 0 {
						layers[next] = layer + 1
						queue.PushBack(next)
					} else if layers[next] == layer {
						return -1
					}
				}
			}
		}
		return layer
	}
	var groupCounts = make(map[int]int)
	var count, group int
	for node := 0; node < n; node++ {
		count = countGroups(node)
		if count == -1 {
			return -1
		}
		group = find(node)
		groupCounts[group] = max(groupCounts[group], count)
	}
	var groups = 0
	for _, count = range groupCounts {
		groups += count
	}
	return groups
}
