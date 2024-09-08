package trees

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/delete-nodes-and-return-forest/

// time: O(N)
// space: O(N)
func delNodes(root *design.TreeNode, toDelete []int) []*design.TreeNode {
	var inDelMap = make(map[int]bool)
	for _, val := range toDelete {
		inDelMap[val] = true
	}
	var dfs func(node *design.TreeNode, isRoot bool) ([]*design.TreeNode, *design.TreeNode) // (roots, child)
	dfs = func(node *design.TreeNode, isRoot bool) ([]*design.TreeNode, *design.TreeNode) {
		if node == nil {
			return nil, nil
		}
		var forest = make([]*design.TreeNode, 0)
		if isRoot && !inDelMap[node.Val] {
			forest = append(forest, node)
		}

		var roots []*design.TreeNode
		roots, node.Left = dfs(node.Left, inDelMap[node.Val])
		forest = append(forest, roots...)

		roots, node.Right = dfs(node.Right, inDelMap[node.Val])
		forest = append(forest, roots...)

		if inDelMap[node.Val] {
			return forest, nil
		}
		return forest, node
	}
	var forest, _ = dfs(root, true)
	return forest
}

// // time: O(N)
// // space: O(N)
// func delNodes(root *design.TreeNode, toDelete []int) []*design.TreeNode {
// 	var inDelMap = make(map[int]bool)
// 	for _, val := range toDelete {
// 		inDelMap[val] = true
// 	}
// 	var dfs func(node *design.TreeNode, isRoot bool) []*design.TreeNode
// 	dfs = func(node *design.TreeNode, isRoot bool) []*design.TreeNode {
// 		var forest = make([]*design.TreeNode, 0)
// 		if isRoot && !inDelMap[node.Val] {
// 			forest = append(forest, node)
// 		}
// 		if node.Left != nil {
// 			forest = append(forest, dfs(node.Left, inDelMap[node.Val])...)
// 			if inDelMap[node.Left.Val] {
// 				node.Left = nil
// 			}
// 		}
// 		if node.Right != nil {
// 			forest = append(forest, dfs(node.Right, inDelMap[node.Val])...)
// 			if inDelMap[node.Right.Val] {
// 				node.Right = nil
// 			}
// 		}
// 		return forest
// 	}
// 	return dfs(root, true)
// }
