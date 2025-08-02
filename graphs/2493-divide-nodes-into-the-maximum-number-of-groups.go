package graphs

import "container/list"

// https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/

// Approach: Union-Find + Breadth-First Search
// Time: O(v*(v+e)), v=# of vertices, e=# of edges
// Space: O(v+e)
func magnificentSets(n int, edges [][]int) int {
	parent := make([]int, n)
	depth := make([]int, n)

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
	union := func(x int, y int) {
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
	graph := make([][]int, n)

	var a, b int
	for idx := range edges {
		a, b = edges[idx][0]-1, edges[idx][1]-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		union(a, b) // groupping connected components
	}

	var (
		curr int
		next int
		ok   bool
	)

	countGroups := func(src int) int {
		queue := list.New()
		queue.PushBack(src)

		layers := make([]int, n)
		layers[src] = 1

		layer := 0
		for size := queue.Len(); size > 0; size = queue.Len() {
			layer++

			for ; size > 0; size-- {
				curr, ok = queue.Remove(queue.Front()).(int)
				_ = ok

				for _, next = range graph[curr] {
					switch layers[next] {
					case 0:
						layers[next] = layer + 1

						queue.PushBack(next)
					case layer:
						return -1
					}
				}
			}
		}

		return layer
	}
	groupCounts := make(map[int]int)

	var count, group int
	for node := 0; node < n; node++ {
		count = countGroups(node)
		if count == -1 {
			return -1
		}

		group = find(node)
		groupCounts[group] = max(groupCounts[group], count)
	}

	groups := 0
	for _, count = range groupCounts {
		groups += count
	}

	return groups
}
