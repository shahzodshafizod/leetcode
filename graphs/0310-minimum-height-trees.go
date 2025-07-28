package graphs

// https://leetcode.com/problems/minimum-height-trees/

// BFS
// time: O(V+E)
// space: O(V*E)
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adjList := make([][]int, n)

	var a, b int
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}

	count := make([]int, n)
	leaves := make([]int, 0)

	var length int
	for node, neighbors := range adjList {
		length = len(neighbors)
		if length == 1 {
			leaves = append(leaves, node)
		}

		count[node] = length
	}
	// hint: delete those nodes with only one connection repeatly,
	// the only 1 or 2 nodes left are the answer.
	for length = len(leaves); n > 2; length = len(leaves) {
		n -= length

		for idx := 0; idx < length; idx++ {
			for _, neighbor := range adjList[leaves[idx]] {
				count[neighbor]--
				if count[neighbor] == 1 {
					leaves = append(leaves, neighbor)
				}
			}
		}

		leaves = leaves[length:]
	}

	return leaves
}

// // DFS: Time Limit Exceeded
// // time: O(V*E)
// // space: O(V*E)
// func findMinHeightTrees(n int, edges [][]int) []int {
// 	if n == 1 {
// 		return []int{0}
// 	}
// 	var adjList = make([][]int, n)
// 	var a, b int
// 	for _, edge := range edges {
// 		a, b = edge[0], edge[1]
// 		adjList[a] = append(adjList[a], b)
// 		adjList[b] = append(adjList[b], a)
// 	}
// 	var dfs func(node int, seen map[int]bool) int
// 	dfs = func(node int, seen map[int]bool) int {
// 		if seen[node] || adjList[node] == nil {
// 			return 0
// 		}
// 		seen[node] = true
// 		var maxheight = -1
// 		var height int
// 		for _, next := range adjList[node] {
// 			height = 1 + dfs(next, seen)
// 			maxheight = max(maxheight, height)
// 		}
// 		return maxheight
// 	}
// 	var roots []int
// 	var minheight = n
// 	for root := range adjList {
// 		height := dfs(root, make(map[int]bool))
// 		if height < minheight {
// 			minheight = height
// 			roots = []int{root}
// 		} else if minheight == height {
// 			roots = append(roots, root)
// 		}
// 	}
// 	return roots
// }
