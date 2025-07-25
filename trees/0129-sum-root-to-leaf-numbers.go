package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/sum-root-to-leaf-numbers/

func sumNumbers(root *pkg.TreeNode) int {
	var dfs func(node *pkg.TreeNode, number int) int
	dfs = func(node *pkg.TreeNode, number int) int {
		number = number*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return number
		}
		sum := 0
		if node.Left != nil {
			sum += dfs(node.Left, number)
		}
		if node.Right != nil {
			sum += dfs(node.Right, number)
		}
		return sum
	}
	return dfs(root, 0)
}
