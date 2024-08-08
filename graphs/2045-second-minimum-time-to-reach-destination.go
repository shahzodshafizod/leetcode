package graphs

// https://leetcode.com/problems/second-minimum-time-to-reach-destination/

// Approach #2: Breadth First Search
// Time: O(N+E)
// Space: O(N+E)
func secondMinimum(n int, edges [][]int, time int, change int) int {
	var adjList = make([][]int, n+1)
	var dist1 = make([]int, n+1)
	var dist2 = make([]int, n+1)
	for idx := 0; idx <= n; idx++ {
		adjList[idx] = make([]int, 0)
		dist1[idx] = -1
		dist2[idx] = -1
	}
	var u, v int
	for _, edge := range edges {
		u, v = edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}
	var queue = [][2]int{{1, 1}} // {city, freq}
	var idx = 0
	dist1[1] = 0
	var timeTaken, node, freq int
	for idx < len(queue) {
		node, freq = queue[idx][0], queue[idx][1]
		idx++
		if freq == 1 {
			timeTaken = dist1[node]
		} else {
			timeTaken = dist2[node]
		}
		if timeTaken/change%2 != 0 {
			timeTaken += change - timeTaken%change
		}
		timeTaken += time
		for _, next := range adjList[node] {
			if dist1[next] == -1 {
				dist1[next] = timeTaken
				queue = append(queue, [2]int{next, 1})
			} else if dist2[next] == -1 && timeTaken != dist1[next] {
				if next == n {
					return timeTaken
				}
				dist2[next] = timeTaken
				queue = append(queue, [2]int{next, 2})
			}
		}
	}
	return 0
}

// // Approach #1: Modified Dijkstra'a Algorithm
// // Time: O(N + E Log N)
// // Space: O(N + E)
// func secondMinimum(n int, edges [][]int, time int, change int) int {
// 	var adjList = make([][]int, n+1)
// 	for idx := range adjList {
// 		adjList[idx] = make([]int, 0)
// 	}
// 	var u, v int
// 	for _, edge := range edges {
// 		u, v = edge[0], edge[1]
// 		adjList[u] = append(adjList[u], v)
// 		adjList[v] = append(adjList[v], u)
// 	}
// 	var pq = design.NewPQ(make([][2]int, 0), func(x, y [2]int) bool { return x[1] > y[1] })
// 	var freq = make([]int, n+1)
// 	var dist1 = make([]int, n+1)
// 	var dist2 = make([]int, n+1)
// 	for idx := range dist1 {
// 		dist1[idx] = math.MaxInt
// 		dist2[idx] = math.MaxInt
// 	}
// 	pq.Push([2]int{1, 0}) // {city, time}
// 	for pq.Len() > 0 {
// 		var city = pq.Pop()
// 		var node, timeTaken = city[0], city[1]
// 		freq[node]++
// 		if node == n && freq[n] == 2 {
// 			return timeTaken
// 		}
// 		if timeTaken/change%2 != 0 {
// 			timeTaken += change - timeTaken%change
// 		}
// 		timeTaken += time
// 		for _, next := range adjList[node] {
// 			if freq[next] == 2 {
// 				continue
// 			}
// 			if timeTaken < dist1[next] {
// 				dist2[next] = dist1[next]
// 				dist1[next] = timeTaken
// 				pq.Push([2]int{next, timeTaken})
// 			} else if timeTaken < dist2[next] && timeTaken != dist1[next] {
// 				dist2[next] = timeTaken
// 				pq.Push([2]int{next, timeTaken})
// 			}
// 		}
// 	}
// 	return 0
// }
