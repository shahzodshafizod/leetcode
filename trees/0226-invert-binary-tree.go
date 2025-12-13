package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/invert-binary-tree/

// Approach #2: Iterative BFS
// Time: O(n) - visit each node once
// Space: O(w) - queue size, where w is maximum width
func invertTree(root *pkg.TreeNode) *pkg.TreeNode {
	if root == nil {
		return nil
	}

	queue := []*pkg.TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Swap left and right children
		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// // Approach #1: Recursive DFS
// // Time: O(n) - visit each node once
// // Space: O(h) - recursion stack, where h is height
// func invertTree(root *pkg.TreeNode) *pkg.TreeNode {
// 	if root != nil {
// 		root.Left, root.Right = root.Right, root.Left
// 		invertTree(root.Left)
// 		invertTree(root.Right)
// 	}

// 	return root
// }
