package trees

import (
	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/

// Approach: Depth-First Search
// Time: O(n+m), n=# of tree nodes, m=len(queries)
// Space: O(n+m)
func treeQueries(root *pkg.TreeNode, queries []int) []int {
	height := make(map[int]int)
	maxHeight := 0

	var dfs func(node *pkg.TreeNode, level int)

	dfs = func(node *pkg.TreeNode, level int) {
		if node == nil {
			return
		}

		height[node.Val] = max(height[node.Val], maxHeight)
		maxHeight = max(maxHeight, level)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
		node.Left, node.Right = node.Right, node.Left
	}
	dfs(root, 0)

	maxHeight = 0

	dfs(root, 0)

	answer := make([]int, 0, len(queries))
	for _, query := range queries {
		answer = append(answer, height[query])
	}

	return answer
}

// // Approach: DFS + BFS
// // Time: O(n+m), n=# of tree nodes, m=len(queries)
// // Space: O(n+m)
// func treeQueries(root *pkg.TreeNode, queries []int) []int {
// 	var height = make(map[int]int)
// 	var dfs func(node *pkg.TreeNode) int
// 	dfs = func(node *pkg.TreeNode) int {
// 		if node == nil {
// 			return -1
// 		}
// 		height[node.Val] = 1 + max(
// 			dfs(node.Left),
// 			dfs(node.Right),
// 		)
// 		return height[node.Val]
// 	}
// 	var bfs = func() {
// 		var queue = []*pkg.TreeNode{root}
// 		var level, max1, max2 = 0, -1, -1
// 		var nextheight, nextmax1, nextmax2 int
// 		for size := len(queue); size > 0; size = len(queue) {
// 			nextmax1, nextmax2 = level, level
// 			for idx := 0; idx < size; idx++ {
// 				node := queue[idx]
// 				height[node.Val] += level
// 				if height[node.Val] == max1 {
// 					height[node.Val] = max2
// 				} else {
// 					height[node.Val] = max1
// 				}
// 				for _, next := range []*pkg.TreeNode{node.Left, node.Right} {
// 					if next != nil {
// 						queue = append(queue, next)
// 						nextheight = level + 1 + height[next.Val]
// 						if nextheight > nextmax1 {
// 							nextmax2, nextmax1 = nextmax1, nextheight
// 						} else if nextheight > nextmax2 {
// 							nextmax2 = nextheight
// 						}
// 					}
// 				}
// 			}
// 			queue = queue[size:]
// 			max1, max2 = nextmax1, nextmax2
// 			level++
// 		}
// 	}
// 	dfs(root) // set down-side height
// 	bfs()     // set tree height if remove current node
// 	var answer = make([]int, 0, len(queries))
// 	for _, query := range queries {
// 		answer = append(answer, height[query])
// 	}
// 	return answer
// }

// // Approach: Brute-Force
// // Time: O(nm), n=# of tree nodes, m=len(queries)
// // Space: O(nm)
// func treeQueries(root *pkg.TreeNode, queries []int) []int {
// 	var getHeight func(node *pkg.TreeNode, level int, skip int) int
// 	getHeight = func(node *pkg.TreeNode, level int, skip int) int {
// 		if node == nil {
// 			return level
// 		}
// 		if node.Val == skip {
// 			return level - 1
// 		}
// 		return max(
// 			getHeight(node.Left, level+1, skip),
// 			getHeight(node.Right, level+1, skip),
// 		)
// 	}
// 	var answer = make([]int, 0, len(queries))
// 	for _, query := range queries {
// 		answer = append(answer, getHeight(root, -1, query))
// 	}
// 	return answer
// }
