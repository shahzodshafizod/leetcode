package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/add-one-row-to-tree/

func addOneRow(root *pkg.TreeNode, val int, depth int) *pkg.TreeNode {
	var dfs func(node *pkg.TreeNode, val int, depth int)

	dfs = func(node *pkg.TreeNode, val int, depth int) {
		if depth == 1 {
			node.Left = &pkg.TreeNode{Val: val, Left: node.Left}
			node.Right = &pkg.TreeNode{Val: val, Right: node.Right}

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
		root = &pkg.TreeNode{Val: val, Left: root}
	} else {
		dfs(root, val, depth-1)
	}

	return root
}
