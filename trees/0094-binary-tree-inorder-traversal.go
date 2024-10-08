package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/binary-tree-inorder-traversal/

// Approach: Morris Traversal
// time: O(n)
// space: O(1)
func inorderTraversal(root *pkg.TreeNode) []int {
	var result = make([]int, 0)
	var curr = root
	var prev *pkg.TreeNode
	for curr != nil {
		if curr.Left == nil {
			result = append(result, curr.Val)
			curr = curr.Right
		} else {
			// find the rightmost node in the left subtree
			prev = curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}
			if prev.Right == curr {
				prev.Right = nil
				result = append(result, curr.Val)
				curr = curr.Right
			} else {
				prev.Right = curr
				curr = curr.Left
			}
		}
	}
	return result
}

// // Approach: Depth-First Search
// // time: O(n)
// // space: O(1)
// func inorderTraversal(root *pkg.TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	var values = inorderTraversal(root.Left)
// 	values = append(values, root.Val)
// 	values = append(values, inorderTraversal(root.Right)...)
// 	return values
// }

// // Follow up: Recursive solution is trivial, could you do it iteratively?
// func inorderTraversal(root *pkg.TreeNode) []int {
// 	var stack = design.NewStack[*pkg.TreeNode]()
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
// 		stack.Push(&pkg.TreeNode{Val: curr.Val})
// 		stack.Push(curr.Left)
// 	}
// 	return values
// }
