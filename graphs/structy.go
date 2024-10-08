package graphs

import (
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

/*
Graph Basics
Graph Algorithms Course for Technical Interviews
https://www.youtube.com/watch?v=2_Uuixtc5i0
(from Alvin Zablan)

graph = nodes(vertices) + edges
nodes: things, edges: relationships
edges are connections between nodes

directed & undirected graphs

(0)---->(2)
 |       |
 |       |
 v       v
(1)<----(4)
 |
 |
 v
(3)---->(5)

how to represent graph information:
- adjacency list
	{
		0: [1, 2],
		1: [3],
		2: [4],
		3: [5],
		4: [1],
		5: []
	}
- matrix

Traversal Approaches (Algorithms):
- Depth-First (uses a Stack)
	0, 1, 3, 5, 2, 4
- Breadth-First (uses a Queue)
	0, 1, 2, 3, 4, 5

asyclic = no cycles
*/

type Structy interface {
	/*
		https://structy.net/problems/has-path

		Write a function, has_path, that takes in a dictionary representing
		the adjacency list of a directed acyclic graph and two nodes (src, dst).
		The function should return a boolean indicating whether or not
		there exists a directed path between the source and destination nodes.
	*/
	HasPath(graph map[byte][]byte, src byte, dts byte) bool
	/*
		https://structy.net/problems/undirected-path

		Write a function, undirected_path, that takes in a list of edges
		for an undirected graph and two nodes (node_A, node_B).
		The function should return a boolean indicating whether or not
		there exists a path between node_A and node_B.
	*/
	UndirectedPath(edges [][]byte, nodeA byte, nodeB byte) bool
	/*
		https://structy.net/problems/connected-components-count

		Write a function, connected_components_count, that takes in
		the adjacency list of an undirected graph. The function should return
		the number of connected components within the graph.
	*/
	ConnectedComponentsCount(graph map[int][]int) int
	/*
		https://structy.net/problems/largest-component

		Write a function, largest_component, that takes in
		the adjacency list of an undirected graph. The function should
		return the size of the largest connected component in the graph.
	*/
	LargestComponent(graph map[int][]int) int
	/*
		https://structy.net/problems/shortest-path

		Write a function, shortest_path, that takes in a list of edges
		for an undirected graph and two nodes (node_A, node_B).
		The function should return the length of the shortest path between A and B.
		Consider the length as the number of edges in the path, not the number of nodes.
		If there is no path between A and B, then return -1.
		You can assume that A and B exist as nodes in the graph.
	*/
	ShortestPath(edges [][]byte, nodeA byte, nodeB byte) int
	/*
		https://structy.net/problems/island-count

		Write a function, island_count, that takes in a grid containing Ws and Ls.
		W represents water and L represents land.
		The function should return the number of islands on the grid.
		An island is a vertically or horizontally connected region of land.
	*/
	IslandCount(grid [][]byte) int
	/*
		https://structy.net/problems/minimum-island

		Write a function, minimum_island, that takes in a grid containing Ws and Ls.
		W represents water and L represents land.
		The function should return the size of the smallest island.
		An island is a vertically or horizontally connected region of land.
		You may assume that the grid contains at least one island.
	*/
	MinimumIsland(grid [][]byte) int
}

type structy struct{}

var _ Structy = &structy{}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (s *structy) HasPath(graph map[byte][]byte, src byte, dst byte) bool {
	var dfs func(src byte, dst byte, visited map[byte]bool) bool
	dfs = func(src byte, dst byte, visited map[byte]bool) bool {
		if src == dst {
			return true
		}
		visited[src] = true
		for _, neighbor := range graph[src] {
			if !visited[neighbor] && dfs(neighbor, dst, visited) {
				return true
			}
		}
		return false
	}
	return dfs(src, dst, make(map[byte]bool))
}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (s *structy) UndirectedPath(edges [][]byte, nodeA byte, nodeB byte) bool {
	var adjList = make(map[byte][]byte)
	var a, b byte
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}
	return s.HasPath(adjList, nodeA, nodeB)
}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (s *structy) ConnectedComponentsCount(graph map[int][]int) int {
	var dfs func(node int, visited map[int]bool)
	dfs = func(node int, visited map[int]bool) {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor, visited)
			}
		}
	}
	var visited = make(map[int]bool)
	var count = 0
	for src := range graph {
		if !visited[src] {
			count++
			dfs(src, visited)
		}
	}
	return count
}

// N = # nodes
// E = # edges
// Time: O(E) = O(N^2)
// Space: O(N)
func (s *structy) LargestComponent(graph map[int][]int) int {
	var count func(node int, visited map[int]bool) int
	count = func(node int, visited map[int]bool) int {
		if visited[node] {
			return 0
		}
		visited[node] = true
		var size = 1
		for _, neighbor := range graph[node] {
			size += count(neighbor, visited)
		}
		return size
	}
	var visited = make(map[int]bool)
	var maxsize = 0
	for src := range graph {
		maxsize = max(maxsize, count(src, visited))
	}
	return maxsize
}

// E = # edges
// Time: O(E)
// Space: O(E)
func (s *structy) ShortestPath(edges [][]byte, nodeA byte, nodeB byte) int {
	var adjList = make(map[byte][]byte)
	var a, b byte
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}
	type node struct {
		val      byte
		distance int
	}
	// BFS
	var queue = pkg.NewQueue[*node]()
	var visited = make(map[byte]bool)
	queue.Enqueue(&node{val: nodeA, distance: 0})
	for !queue.Empty() {
		current := queue.Dequeue()
		visited[current.val] = true
		if current.val == nodeB {
			return current.distance
		}
		for _, neighbor := range adjList[current.val] {
			if !visited[neighbor] {
				queue.Enqueue(&node{val: neighbor, distance: current.distance + 1})
			}
		}
	}
	return -1
}

// r = # rows
// c = # cols
// Time: O(rc)
// Space: O(rc)
func (s *structy) IslandCount(grid [][]byte) int {
	var m, n = len(grid), len(grid[0])
	var directions = [5]int{-1, 0, 1, 0, -1}
	var dfs func(row int, col int, visited [][]bool) int
	dfs = func(row int, col int, visited [][]bool) int {
		if visited[row][col] || grid[row][col] != 'L' {
			return 0
		}
		visited[row][col] = true
		var r, c int
		for dir := 1; dir < 5; dir++ {
			r = row + directions[dir-1]
			c = col + directions[dir]
			if min(r, c) >= 0 && r < m && c < n {
				dfs(r, c, visited)
			}
		}
		return 1
	}
	var visited = make([][]bool, m)
	for row := 0; row < m; row++ {
		visited[row] = make([]bool, n)
	}
	var islands = 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			islands += dfs(row, col, visited)
		}
	}
	return islands
}

// r = # rows
// c = # cols
// Time: O(rc)
// Space: O(rc)
func (s *structy) MinimumIsland(grid [][]byte) int {
	var m, n = len(grid), len(grid[0])
	var directions = [5]int{-1, 0, 1, 0, -1}
	var count func(row int, col int, visited [][]bool) int
	count = func(row int, col int, visited [][]bool) int {
		if visited[row][col] || grid[row][col] != 'L' {
			return 0
		}
		visited[row][col] = true
		var size = 1
		var r, c int
		for dir := 1; dir < 5; dir++ {
			r, c = row+directions[dir-1], col+directions[dir]
			if min(r, c) >= 0 && r < m && c < n {
				size += count(r, c, visited)
			}
		}
		return size
	}
	var visited = make([][]bool, m)
	for row := 0; row < m; row++ {
		visited[row] = make([]bool, n)
	}
	var minSize = math.MaxInt
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if count := count(row, col, visited); count > 0 {
				minSize = min(minSize, count)
			}
		}
	}
	return minSize
}
