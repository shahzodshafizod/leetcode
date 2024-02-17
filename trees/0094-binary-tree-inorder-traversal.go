package trees

// https://leetcode.com/problems/binary-tree-inorder-traversal/

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var values = inorderTraversal(root.Left)
	values = append(values, root.Val)
	values = append(values, inorderTraversal(root.Right)...)
	return values
}

// // Follow up: Recursive solution is trivial, could you do it iteratively?
// func inorderTraversal(root *TreeNode) []int {
// 	var stack = design.NewStack[*TreeNode]()
// 	if root != nil {
// 		stack.Push(root)
// 	}
// 	var values = make([]int, 0)
// 	for stack.Size() > 0 {
// 		curr := stack.Pop()
// 		if curr.Right != nil {
// 			stack.Push(curr.Right)
// 		}
// 		if curr.Left == nil {
// 			values = append(values, curr.Val)
// 			continue
// 		}
// 		stack.Push(&TreeNode{Val: curr.Val})
// 		stack.Push(curr.Left)
// 	}
// 	return values
// }
