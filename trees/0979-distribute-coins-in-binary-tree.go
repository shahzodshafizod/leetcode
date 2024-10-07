package trees

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/distribute-coins-in-binary-tree/

func distributeCoins(root *design.TreeNode) int {
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	var moves = 0
	var dfs func(node *design.TreeNode) int
	dfs = func(node *design.TreeNode) int {
		var left, right = 0, 0
		if node.Left != nil {
			left = dfs(node.Left)
		}
		if node.Right != nil {
			right = dfs(node.Right)
		}
		moves += abs(left) + abs(right)
		return node.Val - 1 + left + right
	}
	dfs(root)
	return moves
}
