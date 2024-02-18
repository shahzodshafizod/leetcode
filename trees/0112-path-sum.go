package trees

// https://leetcode.com/problems/path-sum/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	isLeaf := true
	for _, child := range []*TreeNode{root.Left, root.Right} {
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
