package trees

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/delete-leaves-with-a-given-value/

func removeLeafNodes(root *design.TreeNode, target int) *design.TreeNode {
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
