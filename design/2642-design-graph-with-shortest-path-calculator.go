package design

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/design-graph-with-shortest-path-calculator/

type Graph struct {
	adj [][][2]int // [from -> [to, cost]]
}

// Time: O(E)
// Space: O(E)
func NewGraph(n int, edges [][]int) Graph {
	g := Graph{adj: make([][][2]int, n)}
	for _, edge := range edges {
		g.AddEdge(edge)
	}
	return g
}

func (g *Graph) AddEdge(edge []int) {
	from, to, cost := edge[0], edge[1], edge[2]
	g.adj[from] = append(g.adj[from], [2]int{to, cost})
}

// Approach: Dijkstra's Algorithm
// Time: O(E log E)
// Space: O(E)
func (g *Graph) ShortestPath(node1 int, node2 int) int {
	pq := pkg.NewHeap(
		[][2]int{{node1, 0}}, // [node, cost]
		func(x, y [2]int) bool { return x[1] < y[1] },
	)
	seen := make([]bool, len(g.adj))
	var node, cost, nextNode, nextCost int
	for pq.Len() > 0 {
		item := heap.Pop(pq).([2]int)
		node, cost = item[0], item[1]
		if node == node2 {
			return cost
		}
		seen[node] = true
		for _, nei := range g.adj[node] {
			nextNode, nextCost = nei[0], nei[1]+cost
			if !seen[nextNode] {
				heap.Push(pq, [2]int{nextNode, nextCost})
			}
		}
	}
	return -1
}

// /**
//  * Your Graph object will be instantiated and called as such:
//  * obj := Constructor(n, edges);
//  * obj.AddEdge(edge);
//  * param_2 := obj.ShortestPath(node1,node2);
//  */
