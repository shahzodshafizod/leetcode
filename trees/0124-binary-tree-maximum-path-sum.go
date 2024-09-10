package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/binary-tree-maximum-path-sum/

// Approach: Depth-First Search
// N: # of tree nodes
// H: tree height
// Time: O(N)
// Space: O(H)
func maxPathSum(root *design.TreeNode) int {
	var dfs func(node *design.TreeNode) (int, int)
	dfs = func(node *design.TreeNode) (int, int) {
		if node == nil {
			return -1000, 0
		}
		var lmax, lpath = dfs(node.Left)
		var rmax, rpath = dfs(node.Right)
		lpath = max(0, lpath)
		rpath = max(0, rpath)
		var currMax = max(lmax, rmax, node.Val+lpath+rpath)
		var currPath = node.Val + max(lpath, rpath)
		return currMax, currPath
	}
	maxPath, _ := dfs(root)
	return maxPath
}
