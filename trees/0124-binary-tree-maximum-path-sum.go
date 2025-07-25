package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/binary-tree-maximum-path-sum/

// Approach: Depth-First Search
// N: # of tree nodes
// H: tree height
// Time: O(N)
// Space: O(H)
func maxPathSum(root *pkg.TreeNode) int {
	var dfs func(node *pkg.TreeNode) (int, int)
	dfs = func(node *pkg.TreeNode) (int, int) {
		if node == nil {
			return -1000, 0
		}
		lmax, lpath := dfs(node.Left)
		rmax, rpath := dfs(node.Right)
		lpath = max(0, lpath)
		rpath = max(0, rpath)
		currMax := max(lmax, rmax, node.Val+lpath+rpath)
		currPath := node.Val + max(lpath, rpath)
		return currMax, currPath
	}
	maxPath, _ := dfs(root)
	return maxPath
}
