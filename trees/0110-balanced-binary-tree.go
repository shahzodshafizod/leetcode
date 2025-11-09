package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/balanced-binary-tree/

func isBalanced(root *pkg.TreeNode) bool {
	var dfs func(node *pkg.TreeNode) int

	dfs = func(node *pkg.TreeNode) int {
		if node == nil {
			return 0
		}

		right := dfs(node.Right)
		if right == -1 {
			return -1
		}

		left := dfs(node.Left)
		if left == -1 {
			return -1
		}

		if left-right > 1 || right-left > 1 {
			return -1
		}

		return 1 + max(left, right)
	}

	return dfs(root) != -1
}
