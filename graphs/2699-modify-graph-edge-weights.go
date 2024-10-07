package graphs

import (
	"math"

	"github.com/shahzodshafizod/leetcode/design"
)

// https://leetcode.com/problems/modify-graph-edge-weights/

// Time: O((E+V)logV)
// Space: O(E+V)
func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	// 1. Graph Representation
	var adjList = make([][][2]int, n)
	var u, v int
	for idx := range edges {
		u, v = edges[idx][0], edges[idx][1]
		adjList[u] = append(adjList[u], [2]int{v, idx})
		adjList[v] = append(adjList[v], [2]int{u, idx})
	}

	// 2. Distance Arrays
	var distances = make([][2]int, n)
	for idx := range distances {
		distances[idx][0] = math.MaxInt
		distances[idx][1] = math.MaxInt
	}

	var runDijkstra = func(index int, difference int) {
		var pq = design.NewPQ(
			[][2]int{{source, 0}},
			func(x, y [2]int) bool {
				return x[1] > y[1]
			},
		)
		distances[source][index] = 0
		var currNode, currDist, nextNode, egdeIndex, nextWeight int
		for pq.Len() > 0 {
			var curr = pq.Pop()
			currNode, currDist = curr[0], curr[1]
			if currDist > distances[currNode][index] {
				continue
			}
			for _, next := range adjList[currNode] {
				nextNode, egdeIndex = next[0], next[1]
				nextWeight = edges[egdeIndex][2]
				if nextWeight == -1 {
					nextWeight = 1
				}
				if index == 1 && edges[egdeIndex][2] == -1 {
					var newWeight = difference + distances[nextNode][0] - distances[currNode][1]
					if newWeight > nextWeight {
						nextWeight = newWeight
						edges[egdeIndex][2] = newWeight
					}
				}
				if distances[currNode][index]+nextWeight < distances[nextNode][index] {
					distances[nextNode][index] = distances[currNode][index] + nextWeight
					pq.Push([2]int{nextNode, distances[nextNode][index]})
				}
			}
		}
	}

	// 3. First Dijkstra Run
	runDijkstra(0, 0)

	// 4. Check Feasibility
	var diffirence = target - distances[destination][0]
	if diffirence < 0 {
		return [][]int{}
	}

	// 5. Second Dijkstra Run
	runDijkstra(1, diffirence)

	// 6. Final Check
	if distances[destination][1] < target {
		return [][]int{}
	}

	// 7. Update Edges
	for idx := range edges {
		if edges[idx][2] == -1 {
			edges[idx][2] = 1
		}
	}

	return edges
}
