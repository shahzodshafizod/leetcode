package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/path-sum/

func hasPathSum(root *design.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	isLeaf := true
	for _, child := range []*design.TreeNode{root.Left, root.Right} {
		if child == nil {
			continue
		}
		isLeaf = false
		if hasPathSum(child, targetSum) {
			return true
		}
	}
	return targetSum == 0 && isLeaf
}
