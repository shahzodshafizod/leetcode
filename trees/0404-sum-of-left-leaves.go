package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/sum-of-left-leaves/

// DFS
func sumOfLeftLeaves(root *pkg.TreeNode) int {
	var dfs func(node *pkg.TreeNode, isLeft bool) int
	dfs = func(node *pkg.TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		sum := 0
		if isLeft && node.Left == nil && node.Right == nil {
			sum += node.Val
		}
		return sum + dfs(node.Left, true) + dfs(node.Right, false)
	}
	return dfs(root, false)
}

// // BFS
// func sumOfLeftLeaves(root *pkg.TreeNode) int {
// 	var sum = 0
// 	var queue = []*pkg.TreeNode{root}
// 	for length := len(queue); length > 0; length = len(queue) {
// 		for idx := 0; idx < length; idx++ {
// 			var node = queue[idx]
// 			if node.Left != nil {
// 				if node.Left.Left == nil && node.Left.Right == nil {
// 					sum += node.Left.Val
// 				} else {
// 					queue = append(queue, node.Left)
// 				}
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 		queue = queue[length:]
// 	}
// 	return sum
// }
