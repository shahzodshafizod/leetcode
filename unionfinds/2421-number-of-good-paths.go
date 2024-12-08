package unionfinds

import "sort"

// https://leetcode.com/problems/number-of-good-paths/

// Approach: Union Find
// Time: O(nlogn)
// Space: O(n)
func numberOfGoodPaths(vals []int, edges [][]int) int {
	var n = len(vals)
	var parent = make([]int, n)
	var count = make([]int, n)
	for idx := range parent {
		parent[idx] = idx
		count[idx] = 1
	}
	var find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	sort.Slice(edges, func(i, j int) bool {
		return max(vals[edges[i][0]], vals[edges[i][1]]) < max(vals[edges[j][0]], vals[edges[j][1]])
	})
	var px, py int
	var result = n
	for _, edge := range edges {
		px = find(edge[0])
		py = find(edge[1])
		if vals[px] < vals[py] {
			parent[px] = py
		} else if vals[py] < vals[px] {
			parent[py] = px
		} else {
			parent[py] = px
			result += count[px] * count[py]
			count[px] += count[py]
		}
	}
	return result
}

// // Approach: Brute-Force with Union Find (Time Limit Exceeded)
// // Time: O(?)
// // Space: O(?)
// func numberOfGoodPaths(vals []int, edges [][]int) int {
// 	var n = len(vals)
// 	var parent = make([]int, n)
// 	var nodes = make([]int, n)
// 	for node := range parent {
// 		parent[node] = node
// 		nodes[node] = node
// 	}
// 	var find = func(x int) int {
// 		for x != parent[x] {
// 			parent[x] = parent[parent[x]]
// 			x = parent[x]
// 		}
// 		return x
// 	}
// 	var union = func(x int, y int) {
// 		parent[find(x)] = find(y)
// 	}

// 	var adjList = make([][]int, n)
// 	var node1, node2 int
// 	for idx := range edges {
// 		node1, node2 = edges[idx][0], edges[idx][1]
// 		adjList[node1] = append(adjList[node1], node2)
// 		adjList[node2] = append(adjList[node2], node1)
// 	}
// 	sort.Slice(nodes, func(i, j int) bool { return vals[nodes[i]] < vals[nodes[j]] })
// 	var same = make(map[int]map[int]struct{})
// 	var result = n
// 	for _, node := range nodes {
// 		for _, neighbor := range adjList[node] {
// 			if vals[neighbor] <= vals[node] {
// 				union(neighbor, node)
// 			}
// 		}
// 		for twin := range same[vals[node]] {
// 			if find(twin) == find(node) {
// 				result++
// 			}
// 		}
// 		if same[vals[node]] == nil {
// 			same[vals[node]] = make(map[int]struct{})
// 		}
// 		same[vals[node]][node] = struct{}{}
// 	}
// 	return result
// }
