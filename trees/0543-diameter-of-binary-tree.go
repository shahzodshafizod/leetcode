package trees

import "github.com/shahzodshafizod/leetcode/pkg"

/*
         +----------(1)
     +--(2)--+
 +--(3)     (4)--+
(5)             (6)

diameter: 4 [5,3,2,4,6]
*/

// https://leetcode.com/problems/diameter-of-binary-tree/

func diameterOfBinaryTree(root *pkg.TreeNode) int {
	diameter := 0
	var dfs func(*pkg.TreeNode) int
	dfs = func(node *pkg.TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		diameter = max(diameter, left+right)
		return 1 + max(left, right)
	}
	dfs(root)
	return diameter
}
