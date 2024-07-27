package graphs

// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/

// Approach: Floyd-Warshall Algorithm
// Time: O(N^3)
// Space: O(N^2)
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	var adjMatrix = make([][]int, n)
	for from := range adjMatrix {
		adjMatrix[from] = make([]int, n)
		for to := range adjMatrix[from] {
			adjMatrix[from][to] = 1e9
		}
		adjMatrix[from][from] = 0
	}
	var from, to, weight int
	for _, edge := range edges {
		from, to, weight = edge[0], edge[1], edge[2]
		adjMatrix[from][to] = weight
		adjMatrix[to][from] = weight
	}
	// Floyd-Warshall algorithm to compute
	// shortest paths between all pairs of cities
	for through := 0; through < n; through++ {
		for from := 0; from < n; from++ {
			for to := 0; to < n; to++ {
				adjMatrix[from][to] = min(
					adjMatrix[from][to],
					adjMatrix[from][through]+adjMatrix[through][to], // !!!dangerous section: MAX+MAX becomes negative!!!
				)
			}
		}
	}
	// Determine the city with the fewest number of
	// reachable cities within the adjMatrix threshold
	var city, minConnections = 0, n
	var connections int
	for from := n - 1; from >= 0; from-- {
		connections = 0
		for to := 0; to < n; to++ {
			if from != to && adjMatrix[from][to] <= distanceThreshold {
				connections++
			}
		}
		if connections < minConnections {
			minConnections = connections
			city = from
		}
	}
	return city
}

// // Approach: Dijkstra's Algorithm
// // Time: O(N^3 * Log N)
// // Space: O(N)
// func findTheCity(n int, edges [][]int, distanceThreshold int) int {
// 	var (
// 		from, to, weight int
// 		curr             [2]int
// 		node, dist       int
// 	)
// 	var adjList = make(map[int][][2]int)
// 	for _, edge := range edges {
// 		from, to, weight = edge[0], edge[1], edge[2]
// 		adjList[from] = append(adjList[from], [2]int{to, weight})
// 		adjList[to] = append(adjList[to], [2]int{from, weight})
// 	}

// 	var dijkstra = func(src int) int {
// 		var heap = design.NewPQ(
// 			make([][2]int, 0),
// 			func(x, y [2]int) bool { return x[0] > y[0] },
// 		)
// 		var visited = make(map[int]bool)
// 		heap.Push([2]int{0, src})
// 		for heap.Len() > 0 {
// 			curr = heap.Pop()
// 			dist, node = curr[0], curr[1]
// 			visited[node] = true
// 			for _, next := range adjList[node] {
// 				next, weight := next[0], next[1]+dist
// 				if !visited[next] && weight <= distanceThreshold {
// 					heap.Push([2]int{weight, next})
// 				}
// 			}
// 		}
// 		return len(visited) - 1
// 	}

// 	var isolatedCity, minConnections = -1, n
// 	for src := 0; src < n; src++ {
// 		connections := dijkstra(src)
// 		if connections <= minConnections {
// 			isolatedCity, minConnections = src, connections
// 		}
// 	}
// 	return isolatedCity
// }
