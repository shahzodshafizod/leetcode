package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/evaluate-boolean-binary-tree/

func evaluateTree(root *pkg.TreeNode) bool {
	if root.Left == nil {
		return root.Val == 1
	}
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == 2 {
		return left || right
	}
	return left && right
}
