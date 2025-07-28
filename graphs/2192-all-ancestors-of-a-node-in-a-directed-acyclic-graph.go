package graphs

// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/

// // U N F I N I S H E D
// // Approach #3: Topological Sort (BFS)
// // time: ?
// // space: ?
// func getAncestors(n int, edges [][]int) [][]int {
// 	var adjList = make([][]int, n)
// 	var from, to int
// 	var indegree = make([]int, n)
// 	for _, edge := range edges {
// 		from, to = edge[0], edge[1]
// 		adjList[from] = append(adjList[from], to)
// 		indegree[to]++
// 	}
// 	var queue = make([]int, 0)
// 	for node := 0; node < n; node++ {
// 		if indegree[node] == 0 {
// 			queue = append(queue, node)
// 		}
// 	}
// 	var topoOrder = make([]int, 0, n) // topological order
// 	var start = 0
// 	for start < len(queue) {
// 		var node = queue[start]
// 		start++
// 		topoOrder = append(topoOrder, node)
// 		for _, neighbor := range adjList[node] {
// 			indegree[neighbor]--
// 			if indegree[neighbor] == 0 {
// 				queue = append(queue, neighbor)
// 			}
// 		}
// 	}
// 	fmt.Println("topological order:", topoOrder)
// 	return nil
// }

// Approach 2: Depth First Search (Optimized)
// time: O(N^2+NM) = O(N^2)
// space: O(N+M) = O(N), M: for recursion stack
func getAncestors(n int, edges [][]int) [][]int {
	adjList := make([][]int, n)

	var from, to int
	for _, edge := range edges { // O(E) = O(N^2)
		from, to = edge[0], edge[1]
		adjList[from] = append(adjList[from], to)
	}

	answer := make([][]int, n)

	var dfs func(ancestor int, node int)

	dfs = func(ancestor int, node int) {
		for _, neighbor := range adjList[node] {
			if len(answer[neighbor]) == 0 ||
				answer[neighbor][len(answer[neighbor])-1] != ancestor {
				answer[neighbor] = append(answer[neighbor], ancestor)
				dfs(ancestor, neighbor)
			}
		}
	}
	for ancestor := 0; ancestor < n; ancestor++ { // O(N)
		dfs(ancestor, ancestor) // O(M) = maximal path length
	}

	return answer
}

// // Approach #1: Depth First Search (Reversed Graph)
// // time: O(N^2 + NM)
// // space: O(N+M)
// func getAncestors(n int, edges [][]int) [][]int {
// 	var adjList = make([][]int, n)
// 	var from, to int
// 	// Populate the adjacency list with reversed edges
// 	// because we traverse from the node to find its ancestors
// 	// and ancestors become descendants
// 	for _, edge := range edges { // O(E) = O(N^2)
// 		from, to = edge[0], edge[1]
// 		adjList[to] = append(adjList[to], from)
// 	}
// 	var findChildren func(node int, visited map[int]bool)
// 	findChildren = func(node int, visited map[int]bool) {
// 		if visited[node] {
// 			return
// 		}
// 		visited[node] = true
// 		for _, child := range adjList[node] { // O(M) - maximal path length
// 			findChildren(child, visited)
// 		}
// 	}
// 	var answer = make([][]int, n)
// 	for idx := range answer { // O(N)
// 		// answer[idx] = make([]int, 0)
// 		var visited = make(map[int]bool)
// 		findChildren(idx, visited)        // O(M)
// 		for node := 0; node < n; node++ { // O(N)
// 			if node != idx && visited[node] {
// 				answer[idx] = append(answer[idx], node)
// 			}
// 		}
// 	}
// 	return answer
// }
