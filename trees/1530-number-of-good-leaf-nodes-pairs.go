package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/

// Approach #2: Lowest Common Ancestor (LCA)
// N = # of nodes
// D = distance
// Time: O(N x D^2)
// Space: O(Log N)
func countPairs(root *design.TreeNode, distance int) int {
	var pairs = 0
	var dfs func(node *design.TreeNode, level int) []int
	dfs = func(node *design.TreeNode, level int) []int {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{level}
		}
		var leftLeaves = dfs(node.Left, level+1)
		var rightLeaves = dfs(node.Right, level+1)
		var leaves = make([]int, 0)
		for _, l := range leftLeaves {
			if l-level < distance {
				leaves = append(leaves, l)
			}
			l -= level
			for _, r := range rightLeaves {
				r -= level
				if l+r <= distance {
					pairs++
				}
			}
		}
		for _, r := range rightLeaves {
			if r-level < distance {
				leaves = append(leaves, r)
			}
		}
		return leaves
	}
	dfs(root, 0)
	return pairs
}

// // Approach #1: Graph Conversion
// // Time: O(N^2)
// // Space: O(N)
// func countPairs(root *design.TreeNode, distance int) int {
// 	var adjList = make(map[*design.TreeNode][]*design.TreeNode)
// 	var buildGraph func(parent *design.TreeNode, node *design.TreeNode)
// 	buildGraph = func(parent *design.TreeNode, node *design.TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		if parent != nil {
// 			adjList[node] = append(adjList[node], parent)
// 			adjList[parent] = append(adjList[parent], node)
// 		}
// 		buildGraph(node, node.Left)
// 		buildGraph(node, node.Right)
// 	}
// 	buildGraph(nil, root)

// 	var countPairs func(node *design.TreeNode, dist int, visited map[*design.TreeNode]bool) int
// 	countPairs = func(node *design.TreeNode, dist int, visited map[*design.TreeNode]bool) int {
// 		if visited[node] || dist > distance {
// 			return 0
// 		}
// 		visited[node] = true
// 		var pairs = 0
// 		if node.Left == nil && node.Right == nil {
// 			pairs++
// 		}
// 		for _, neighbor := range adjList[node] {
// 			pairs += countPairs(neighbor, dist+1, visited)
// 		}
// 		return pairs
// 	}

// 	var pairs = 0
// 	for node := range adjList { // O(N)
// 		if node.Left == nil && node.Right == nil {
// 			pairs += countPairs(node, 0, make(map[*design.TreeNode]bool)) - 1 // O(N)
// 		}
// 	}
// 	return pairs / 2 // div into 2 means: we encounter nodes A and B twice: (A,B) and (B,A)
// }
