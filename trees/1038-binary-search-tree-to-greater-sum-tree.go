package trees

import "github.com/shahzodshafizod/leetcode/pkg"

/*
Note: This question is the same as 538:
https://leetcode.com/problems/convert-bst-to-greater-tree/
*/

// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

// Approach #2: Morris Traversal
// time: O(n)
// space: O(1)
func bstToGst(root *pkg.TreeNode) *pkg.TreeNode {
	var sum = 0
	var prev *pkg.TreeNode
	var curr = root
	for curr != nil {
		if curr.Right == nil {
			sum += curr.Val
			curr.Val = sum
			curr = curr.Left
		} else {
			prev = curr.Right
			for prev.Left != nil && prev.Left != curr {
				prev = prev.Left
			}
			if prev.Left == curr {
				prev.Left = nil
				sum += curr.Val
				curr.Val = sum
				curr = curr.Left
			} else {
				prev.Left = curr
				curr = curr.Right
			}
		}
	}
	return root
}

// // Approach #1: Reversed InOrder Depth-First Search Traversal
// // time: O(n)
// // space: O(n) for recursion stack
// func bstToGst(root *pkg.TreeNode) *pkg.TreeNode {
// 	var dfs func(node *pkg.TreeNode, sum int) int
// 	dfs = func(node *pkg.TreeNode, sum int) int {
// 		if node == nil {
// 			return sum
// 		}
// 		node.Val += dfs(node.Right, sum)
// 		return dfs(node.Left, node.Val)
// 	}
// 	dfs(root, 0)
// 	return root
// }
