package trees

import (
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-largest-value-in-each-tree-row/

// Approach: Breadth-First Search
// Time: O(n), n=# of nodes
// Space: O(w), w=width of tree
func largestValues(root *pkg.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	largest := make([]int, 0)
	queue := []*pkg.TreeNode{root}

	var maxval int

	for len(queue) > 0 {
		maxval = math.MinInt
		nextq := make([]*pkg.TreeNode, 0)

		for _, node := range queue {
			maxval = max(maxval, node.Val)

			if node.Left != nil {
				nextq = append(nextq, node.Left)
			}

			if node.Right != nil {
				nextq = append(nextq, node.Right)
			}
		}

		largest = append(largest, maxval)
		queue = nextq
	}

	return largest
}

// // Approach: Depth-First Search
// // Time: O(n), n=# of nodes
// // Space: O(h), h=height of tree
// func largestValues(root *pkg.TreeNode) []int {
// 	var largest = make([]int, 0)
// 	var dfs func(node *pkg.TreeNode, level int)
// 	dfs = func(node *pkg.TreeNode, level int) {
// 		if node == nil {
// 			return
// 		}
// 		if len(largest) == level {
// 			largest = append(largest, node.Val)
// 		} else {
// 			largest[level] = max(largest[level], node.Val)
// 		}
// 		dfs(node.Left, level+1)
// 		dfs(node.Right, level+1)
// 	}
// 	dfs(root, 0)
// 	return largest
// }
