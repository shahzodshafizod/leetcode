package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/path-sum/

func hasPathSum(root *pkg.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	isLeaf := true
	for _, child := range []*pkg.TreeNode{root.Left, root.Right} {
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

// func hasPathSum(root *pkg.TreeNode, targetSum int) bool {
// 	var dfs func(node *pkg.TreeNode, currSum int) bool
// 	dfs = func(node *pkg.TreeNode, currSum int) bool {
// 		if node == nil {
// 			return false
// 		}
// 		currSum += node.Val
// 		if node.Left == nil && node.Right == nil {
// 			return currSum == targetSum
// 		}
// 		return dfs(node.Left, currSum) || dfs(node.Right, currSum)
// 	}
// 	return dfs(root, 0)
// }
