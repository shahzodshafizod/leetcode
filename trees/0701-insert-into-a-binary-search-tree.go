package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/insert-into-a-binary-search-tree/

func insertIntoBST(root *design.TreeNode, val int) *design.TreeNode {
	if root == nil {
		return &design.TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
