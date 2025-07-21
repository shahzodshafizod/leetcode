package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

func constructFromPrePost(preorder []int, postorder []int) *pkg.TreeNode {
	preidx, postidx := 0, 0
	var construct func() *pkg.TreeNode
	construct = func() *pkg.TreeNode {
		node := &pkg.TreeNode{Val: preorder[preidx]}
		preidx++
		if node.Val != postorder[postidx] {
			node.Left = construct()
		}
		if node.Val != postorder[postidx] {
			node.Right = construct()
		}
		postidx++
		return node
	}
	return construct()
}
