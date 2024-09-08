package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/find-bottom-left-tree-value/

// DFS
func findBottomLeftValue(root *design.TreeNode) int {
	var value int
	var maxLevel = -1
	var dfs func(node *design.TreeNode, level int)
	dfs = func(node *design.TreeNode, level int) {
		if node == nil {
			return
		}
		if level > maxLevel {
			maxLevel = level
			value = node.Val
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return value
}

// // BFS
// func findBottomLeftValue(root *design.TreeNode) int {
// 	var value int
// 	var queue = []*design.TreeNode{root}
// 	for length := len(queue); length > 0; length = len(queue) {
// 		value = queue[0].Val
// 		for i := 0; i < length; i++ {
// 			node := queue[i]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 		}
// 		queue = queue[length:]
// 	}
// 	return value
// }
