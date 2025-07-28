package graphs

// https://leetcode.com/problems/find-minimum-diameter-after-merging-two-trees/

// Approach: Depth-First Search
// Time: O(n+m), n=len(edges1), m=len(edges2)
// Space: O(n+m)
func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	var dfs func(adjList [][]int, parent int, node int) (int, int)

	dfs = func(adjList [][]int, parent int, node int) (int, int) {
		endpoint, distance := node, 0

		var point, dist int

		for _, next := range adjList[node] {
			if next == parent {
				continue
			}

			point, dist = dfs(adjList, node, next)
			dist++ // include the edge between node and next

			if dist > distance {
				distance = dist
				endpoint = point
			}
		}

		return endpoint, distance
	}
	findDiameter := func(edges [][]int) int {
		n := len(edges) + 1
		adjList := make([][]int, n)

		var a, b int
		for idx := range edges {
			a, b = edges[idx][0], edges[idx][1]
			adjList[a] = append(adjList[a], b)
			adjList[b] = append(adjList[b], a)
		}

		endpoint, _ := dfs(adjList, -1, 0)
		_, diameter := dfs(adjList, -1, endpoint)

		return diameter
	}
	d1 := findDiameter(edges1)
	d2 := findDiameter(edges2)

	return max(d1, d2, (d1+1)/2+(d2+1)/2+1)
}

// // Approach: Breadth-First Search
// // Time: O(n+m), n=len(edges1), m=len(edges2)
// // Space: O(n+m)
// func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
// 	var bfs = func(adjList [][]int, node int) (int, int) {
// 		var queue = [][3]int{{-1, node, 0}}
// 		var parent, distance int
// 		for idx := 0; idx < len(queue); idx++ {
// 			parent = queue[idx][0]
// 			node = queue[idx][1]
// 			distance = queue[idx][2]
// 			for _, next := range adjList[node] {
// 				if next == parent {
// 					continue
// 				}
// 				queue = append(queue, [3]int{node, next, distance + 1})
// 			}
// 		}
// 		return node, distance
// 	}
// 	var findDiameter = func(edges [][]int) int {
// 		var n = len(edges) + 1
// 		var adjList = make([][]int, n)
// 		var a, b int
// 		for idx := range edges {
// 			a, b = edges[idx][0], edges[idx][1]
// 			adjList[a] = append(adjList[a], b)
// 			adjList[b] = append(adjList[b], a)
// 		}
// 		var endpoint, _ = bfs(adjList, 0)
// 		var _, diameter = bfs(adjList, endpoint)
// 		return diameter
// 	}
// 	var d1 = findDiameter(edges1)
// 	var d2 = findDiameter(edges2)
// 	return max(d1, d2, (d1+1)/2+(d2+1)/2+1)
// }
