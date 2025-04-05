package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/

// Note: This question is the same as 865:
// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/

func lcaDeepestLeaves(root *pkg.TreeNode) *pkg.TreeNode {
	var dfs func(node *pkg.TreeNode) (*pkg.TreeNode, int)
	dfs = func(node *pkg.TreeNode) (*pkg.TreeNode, int) {
		if node == nil {
			return nil, 0
		}
		var lanc, lheight = dfs(node.Left)
		var ranc, rheight = dfs(node.Right)
		if lheight == rheight {
			return node, lheight + 1
		}
		if lheight > rheight {
			return lanc, lheight + 1
		}
		return ranc, rheight + 1
	}
	var anc, _ = dfs(root)
	return anc
}
