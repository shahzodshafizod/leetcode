package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/search-in-a-binary-search-tree/

func searchBST(root *pkg.TreeNode, val int) *pkg.TreeNode {
	for root != nil && root.Val != val {
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}

// func searchBST(root *pkg.TreeNode, val int) *pkg.TreeNode {
// 	if root == nil || root.Val == val {
// 		return root
// 	}
// 	if val < root.Val {
// 		return searchBST(root.Left, val)
// 	}
// 	return searchBST(root.Right, val)
// }
