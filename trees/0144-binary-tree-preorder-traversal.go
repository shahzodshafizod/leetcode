package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/binary-tree-preorder-traversal/

func preorderTraversal(root *pkg.TreeNode) []int {
	var dfs func(node *pkg.TreeNode) []int
	dfs = func(node *pkg.TreeNode) []int {
		if node == nil {
			return []int{}
		}
		vals := []int{node.Val}
		vals = append(vals, dfs(node.Left)...)
		vals = append(vals, dfs(node.Right)...)
		return vals
	}
	return dfs(root)
}
