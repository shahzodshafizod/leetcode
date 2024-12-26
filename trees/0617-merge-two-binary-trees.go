package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/merge-two-binary-trees/

// Approach: Breadth-First Search
// Time: O(max(n1, n2)), n1=# of tree1 node, n2=# of tree2 nodes
// Space: O(max(w1, w2)), w1=width of tree1, w2=width of tree2
func mergeTrees(root1 *pkg.TreeNode, root2 *pkg.TreeNode) *pkg.TreeNode {
	if root1 == nil {
		return root2
	}
	var stack = [][2]*pkg.TreeNode{{root1, root2}}
	var size = 1
	var node1, node2 *pkg.TreeNode
	for size > 0 {
		node1, node2 = stack[size-1][0], stack[size-1][1]
		stack = stack[:size-1]
		size--
		if node1 == nil || node2 == nil {
			continue
		}
		node1.Val += node2.Val

		if node1.Left == nil {
			node1.Left = node2.Left
		} else {
			stack = append(stack, [2]*pkg.TreeNode{node1.Left, node2.Left})
			size++
		}
		if node1.Right == nil {
			node1.Right = node2.Right
		} else {
			stack = append(stack, [2]*pkg.TreeNode{node1.Right, node2.Right})
			size++
		}
	}
	return root1
}

// // Approach: Depth-First Search
// // Time: O(max(n1, n2)), n1=# of tree1 node, n2=# of tree2 nodes
// // Space: O(max(h1, h2)), h1=height of tree1, h2=height of tree2
// func mergeTrees(root1 *pkg.TreeNode, root2 *pkg.TreeNode) *pkg.TreeNode {
// 	if root1 == nil {
// 		return root2
// 	}
// 	if root2 == nil {
// 		return root1
// 	}
// 	root1.Val += root2.Val
// 	root1.Left = mergeTrees(root1.Left, root2.Left)
// 	root1.Right = mergeTrees(root1.Right, root2.Right)
// 	return root1
// }
