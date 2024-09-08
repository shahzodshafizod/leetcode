package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/add-one-row-to-tree/

func addOneRow(root *design.TreeNode, val int, depth int) *design.TreeNode {
	var dfs func(node *design.TreeNode, val int, depth int)
	dfs = func(node *design.TreeNode, val int, depth int) {
		if depth == 1 {
			node.Left = &design.TreeNode{Val: val, Left: node.Left}
			node.Right = &design.TreeNode{Val: val, Right: node.Right}
			return
		}
		if node.Left != nil {
			dfs(node.Left, val, depth-1)
		}
		if node.Right != nil {
			dfs(node.Right, val, depth-1)
		}
	}
	if depth == 1 {
		root = &design.TreeNode{Val: val, Left: root}
	} else {
		dfs(root, val, depth-1)
	}
	return root
}
