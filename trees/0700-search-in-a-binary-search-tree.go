package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/search-in-a-binary-search-tree/

func searchBST(root *design.TreeNode, val int) *design.TreeNode {
	for root != nil && root.Val != val {
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}

// func searchBST(root *design.TreeNode, val int) *design.TreeNode {
// 	if root == nil || root.Val == val {
// 		return root
// 	}
// 	if val < root.Val {
// 		return searchBST(root.Left, val)
// 	}
// 	return searchBST(root.Right, val)
// }
