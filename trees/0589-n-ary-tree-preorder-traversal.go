package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

// BFS
func preorder(root *pkg.NTreeNode) []int {
	var values = make([]int, 0)
	var stack = make([]*pkg.NTreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for length := len(stack); length > 0; length = len(stack) {
		var top = stack[length-1]
		stack = stack[:length-1]
		values = append(values, top.Val)
		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack, top.Children[i])
		}
	}
	return values
}

// // DFS
// func preorder(root *pkg.NTreeNode) []int {
// 	var values = make([]int, 0)
// 	if root == nil {
// 		return values
// 	}
// 	values = append(values, root.Val)
// 	for _, child := range root.Children {
// 		values = append(values, preorder(child)...)
// 	}
// 	return values
// }
