package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/symmetric-tree/

// Approach: Recursively
// Time: O(n), n=# of tree nodes
// Space: O(h), h=height of tree
func isSymmetric(root *pkg.TreeNode) bool {
	var isMirror func(node1 *pkg.TreeNode, node2 *pkg.TreeNode) bool

	isMirror = func(node1 *pkg.TreeNode, node2 *pkg.TreeNode) bool {
		if node1 == nil || node2 == nil {
			return node1 == node2
		}

		if node1.Val != node2.Val {
			return false
		}

		return isMirror(node1.Left, node2.Right) &&
			isMirror(node1.Right, node2.Left)
	}

	return isMirror(root.Left, root.Right)
}

// // Approach: Iteratively
// // Time: O(n), n=# of tree nodes
// // Space: O(w), w=width of tree
// func isSymmetric(root *pkg.TreeNode) bool {
// 	var lstack = []*pkg.TreeNode{root.Left}
// 	var rstack = []*pkg.TreeNode{root.Right}
// 	var lsize, rsize = 1, 1
// 	var lnode, rnode *pkg.TreeNode
// 	for lsize != 0 && rsize != 0 {
// 		lnode = lstack[lsize-1]
// 		rnode = rstack[rsize-1]
// 		lstack = lstack[:lsize-1]
// 		rstack = rstack[:rsize-1]
// 		lsize--
// 		rsize--
// 		if lnode == nil && rnode == nil {
// 			continue
// 		}
// 		if lnode == nil || rnode == nil || lnode.Val != rnode.Val {
// 			return false
// 		}
// 		lstack = append(lstack, lnode.Left, lnode.Right)
// 		rstack = append(rstack, rnode.Right, rnode.Left)
// 		lsize += 2
// 		rsize += 2
// 	}
// 	return lsize == 0 && rsize == 0
// }
