package graphs

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

/*
Algorithmic approaches:
	- Sorting
	- Greedy Method
	- Divide and Conquer
	- Dynamic Programming
	- Backtracking

Problem:
There are n network nodes labelled 1 to N.
	e.g.: n=5, nodes=1,2,3,4,5

Given a times array, containing edges represented by arrays [u, v, w]
where u is the source node, v is the target node, and w is the time
taken to travel from the source node to the target node.
	e.g.: times = [[1,2,9],[1,4,2],[2,5,1],[4,2,4],[4,5,6],[3,2,3],[3,1,5],[5,3,7]]

Send a signal from node k, return how long it takes for all nodes
to receive the signal. Return -1 if it's impossible.

Step 1: Verify the constraints
	- Can the graph be unconnected?
		: Yes, account for an unconnected graph.
	- Can the time be negative integers?
		: No, the weight of edges is always positive.
Step 2: Write out some test cases
	- times=[[2,3,4]], n=3, k=2, delay: -1
	- times=[[1,2,8],[3,1,3]], n=3, k=1, delay=-1

Graphs: Dijkstra's Algorithm (Greedy Method)
Dijkstra's Algorithm can only be applied to graphs that are directed and weighted.
Greedy Method is an algorithmic paradigm. It only applies when we're working with
optimization problem (min, max, shortest time).

Bellman-Ford Algorithm (Dynamic Programming)
There shouldn't be a negative cycle.

If there are just positive edges in graph, use Dijkstra's algorithm, else user Bellman-Ford's.
*/

// https://leetcode.com/problems/network-delay-time/

type distance struct {
	target int
	weight int
}

// Dijkstra's Algorithm
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][]*distance)

	for _, time := range times {
		source, target, weight := time[0], time[1], time[2]
		graph[source] = append(graph[source], &distance{target: target, weight: weight})
	}

	minHeap := pkg.NewHeap(
		[]*distance{{target: k, weight: 0}},
		func(x, y *distance) bool { return x.weight < y.weight },
	)
	arrivalTime := make(map[int]int)

	for minHeap.Len() > 0 {
		dist := heap.Pop(minHeap).(*distance)
		if _, found := arrivalTime[dist.target]; found {
			continue
		}

		arrivalTime[dist.target] = dist.weight

		for _, next := range graph[dist.target] {
			next.weight += dist.weight
			heap.Push(minHeap, next)
		}
	}

	if len(arrivalTime) != n {
		return -1
	}

	maxDistance := 0
	for _, time := range arrivalTime {
		if time > maxDistance {
			maxDistance = time
		}
	}

	return maxDistance
}

// // Dijkstra's Algorithm
// func networkDelayTime(times [][]int, n int, k int) int {
// 	var distances = make([]int, n)
// 	var adjList = make([][][2]int, n)
// 	for i := 0; i < n; i++ {
// 		distances[i] = math.MaxInt
// 		adjList[i] = make([][2]int, 0)
// 	}
// 	distances[k-1] = 0
// 	for _, time := range times {
// 		var source = time[0]
// 		var target = time[1]
// 		var weight = time[2]
// 		adjList[source-1] = append(adjList[source-1], [2]int{target - 1, weight})
// 	}
// 	var minHeap = trees.NewPriorityQueue(func(a int, b int) bool { return distances[a] > distances[b] })
// 	minHeap.Push(k - 1)
// 	for !minHeap.IsEmpty() {
// 		var currentVertex = minHeap.Pop()
// 		var adjacent = adjList[currentVertex]
// 		for _, adj := range adjacent {
// 			var nextVertex = adj[0]
// 			var weight = adj[1]
// 			if distances[currentVertex]+weight < distances[nextVertex] {
// 				distances[nextVertex] = distances[currentVertex] + weight
// 				if !minHeap.Has(nextVertex) {
// 					minHeap.Push(nextVertex)
// 				}
// 			}
// 		}
// 	}
// 	var maxDistance = distances[0]
// 	for _, distance := range distances {
// 		maxDistance = max(maxDistance, distance)
// 	}
// 	if maxDistance == math.MaxInt {
// 		maxDistance = -1
// 	}
// 	return maxDistance
// }

// // Bellman-Ford Algorithm
// func networkDelayTime(times [][]int, n int, k int) int {
// 	var distances = make([]int, n)
// 	for i := 0; i < n; i++ {
// 		distances[i] = math.MaxInt
// 	}
// 	distances[k-1] = 0
// 	for i := n - 1; i > 0; i-- {
// 		var changed = false
// 		for _, time := range times {
// 			var source = time[0]
// 			var target = time[1]
// 			var weight = time[2]
// 			if distances[source-1] != math.MaxInt &&
// 				distances[source-1]+weight < distances[target-1] {
// 				distances[target-1] = distances[source-1] + weight
// 				changed = true
// 			}
// 		}
// 		if !changed {
// 			break
// 		}
// 	}
// 	var maxDistance = distances[0]
// 	for _, distance := range distances {
// 		maxDistance = max(maxDistance, distance)
// 	}
// 	if maxDistance == math.MaxInt {
// 		maxDistance = -1
// 	}
// 	return maxDistance
// }
