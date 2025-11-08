package graphs

// https://leetcode.com/problems/longest-special-path/

// Approach 2: Depth-First Search & Sliding Window
// Time: O(n)
// Space: O(n)
func longestSpecialPath(edges [][]int, nums []int) []int {
	n := len(nums)
	graph := make([][][2]int, n)

	var u, v, length int
	for _, edge := range edges {
		u, v, length = edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], [2]int{v, length})
		graph[v] = append(graph[v], [2]int{u, length})
	}

	result := []int{0, 1}

	var costs []int

	last := make(map[int]*int)

	var dfs func(parent int, node int, cost int, start int)

	dfs = func(parent int, node int, cost int, start int) {
		color := last[nums[node]]
		last[nums[node]] = new(int)
		*last[nums[node]] = len(costs)
		costs = append(costs, cost)

		length := cost - costs[start]
		cnt := len(costs) - start

		if length > result[0] {
			result[0] = length
			result[1] = cnt
		} else if length == result[0] && cnt < result[1] {
			result[1] = cnt
		}

		var nei, ncost, nstart int
		for _, edge := range graph[node] {
			nei, ncost = edge[0], edge[1]
			if nei == parent {
				continue
			}

			nstart = start
			if last[nums[nei]] != nil && start <= *last[nums[nei]] {
				nstart = *last[nums[nei]] + 1
			}

			dfs(node, nei, cost+ncost, nstart)
		}

		last[nums[node]] = color
		costs = costs[:len(costs)-1]
	}

	dfs(-1, 0, 0, 0)

	return result
}

// // Approach 1: Brute Force
// // Time: O(nn)
// // Space: O(n)
// func longestSpecialPath(edges [][]int, nums []int) []int {
// 	n := len(nums)
// 	graph := make([][][2]int, n)

// 	for _, edge := range edges {
// 		u, v, length := edge[0], edge[1], edge[2]
// 		graph[u] = append(graph[u], [2]int{v, length})
// 		graph[v] = append(graph[v], [2]int{u, length})
// 	}

// 	downward := make([][][2]int, n)
// 	visitedTree := make([]bool, n)

// 	var buildTree func(node, parent int)

// 	buildTree = func(node, parent int) {
// 		visitedTree[node] = true
// 		for _, edge := range graph[node] {
// 			neighbor, length := edge[0], edge[1]
// 			if neighbor != parent && !visitedTree[neighbor] {
// 				downward[node] = append(downward[node], [2]int{neighbor, length})
// 				buildTree(neighbor, node)
// 			}
// 		}
// 	}

// 	buildTree(0, -1)

// 	maxLength, minCnt := 0, 1

// 	visited := make(map[int]bool)

// 	var dfs func(node int, length, cnt int)

// 	dfs = func(node int, length, cnt int) {
// 		visited[nums[node]] = true

// 		if length > maxLength {
// 			maxLength = length
// 			minCnt = cnt
// 		} else if length == maxLength && cnt < minCnt {
// 			minCnt = cnt
// 		}

// 		for _, edge := range downward[node] {
// 			child, edgeLen := edge[0], edge[1]
// 			if !visited[nums[child]] {
// 				dfs(child, length+edgeLen, cnt+1)
// 			}
// 		}

// 		visited[nums[node]] = false
// 	}

// 	for start := range n {
// 		dfs(start, 0, 1)
// 	}

// 	return []int{maxLength, minCnt}
// }
