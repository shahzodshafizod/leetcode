package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/

// // Iterative
// func postorder(root *pkg.NTreeNode) []int {
// 	var stack = []*pkg.NTreeNode{}
// 	if root != nil {
// 		stack = append(stack, root)
// 	}
// 	var order = []int{}
// 	for len(stack) > 0 {
// 		var node = stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		order = append([]int{node.Val}, order...)
// 		stack = append(stack, node.Children...)
// 	}
// 	return order
// }

// Recursive
func postorder(root *pkg.NTreeNode) []int {
	if root == nil {
		return []int{}
	}

	order := make([]int, 0)
	for _, child := range root.Children {
		order = append(order, postorder(child)...)
	}

	order = append(order, root.Val)

	return order
}
