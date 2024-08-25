package trees

// https://leetcode.com/problems/binary-tree-postorder-traversal/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var order = postorderTraversal(root.Left)
	order = append(order, postorderTraversal(root.Right)...)
	order = append(order, root.Val)
	return order
}
