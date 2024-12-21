package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/minimum-depth-of-binary-tree/

// Approach: Breadth-First Search
// Time: O(n), n=# of nodes
// Space: O(w), w=width of tree
func minDepth(root *pkg.TreeNode) int {
	var queue = make([]*pkg.TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	var depth = 0
	for size := len(queue); size > 0; size = len(queue) {
		depth++
		for idx := 0; idx < size; idx++ {
			node := queue[idx]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return 0
}

// // Approach: Depth-First Search
// // Time: O(n), n=# of nodes
// // Space: O(d), d=depth of tree
// func minDepth(root *pkg.TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	var left = minDepth(root.Left)
// 	var right = minDepth(root.Right)
// 	if left == 0 || right == 0 {
// 		return 1 + left + right
// 	}
// 	return 1 + min(left, right)
// }
