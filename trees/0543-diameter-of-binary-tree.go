package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

/*
         +----------(1)
     +--(2)--+
 +--(3)     (4)--+
(5)             (6)

diameter: 4 [5,3,2,4,6]
*/

// https://leetcode.com/problems/diameter-of-binary-tree/

func diameterOfBinaryTree(root *design.TreeNode) int {
	var diameter = 0
	var dfs func(*design.TreeNode) int
	dfs = func(node *design.TreeNode) int {
		if node == nil {
			return 0
		}
		var left = dfs(node.Left)
		var right = dfs(node.Right)
		diameter = max(diameter, left+right)
		return 1 + max(left, right)
	}
	dfs(root)
	return diameter
}
