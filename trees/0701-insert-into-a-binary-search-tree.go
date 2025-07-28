package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/insert-into-a-binary-search-tree/

func insertIntoBST(root *pkg.TreeNode, val int) *pkg.TreeNode {
	if root == nil {
		return &pkg.TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}
