package trees

// https://leetcode.com/problems/insert-into-a-binary-search-tree/

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
