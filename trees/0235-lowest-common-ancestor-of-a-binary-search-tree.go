package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *pkg.TreeNode) *pkg.TreeNode {
	for (root.Val-p.Val)*(root.Val-q.Val) > 0 {
		if root.Val > p.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// func lowestCommonAncestor(root, p, q *pkg.TreeNode) *pkg.TreeNode {
// 	if root.Val > p.Val && root.Val > q.Val {
// 		return lowestCommonAncestor(root.Left, p, q)
// 	}
// 	if root.Val < p.Val && root.Val < q.Val {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	}
// 	return root
// }
