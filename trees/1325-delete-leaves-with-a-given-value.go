package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/delete-leaves-with-a-given-value/

func removeLeafNodes(root *pkg.TreeNode, target int) *pkg.TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}

	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}
