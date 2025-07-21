package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/delete-nodes-and-return-forest/

// time: O(N)
// space: O(N)
func delNodes(root *pkg.TreeNode, toDelete []int) []*pkg.TreeNode {
	inDelMap := make(map[int]bool)
	for _, val := range toDelete {
		inDelMap[val] = true
	}
	var dfs func(node *pkg.TreeNode, isRoot bool) ([]*pkg.TreeNode, *pkg.TreeNode) // (roots, child)
	dfs = func(node *pkg.TreeNode, isRoot bool) ([]*pkg.TreeNode, *pkg.TreeNode) {
		if node == nil {
			return nil, nil
		}
		forest := make([]*pkg.TreeNode, 0)
		if isRoot && !inDelMap[node.Val] {
			forest = append(forest, node)
		}

		var roots []*pkg.TreeNode
		roots, node.Left = dfs(node.Left, inDelMap[node.Val])
		forest = append(forest, roots...)

		roots, node.Right = dfs(node.Right, inDelMap[node.Val])
		forest = append(forest, roots...)

		if inDelMap[node.Val] {
			return forest, nil
		}
		return forest, node
	}
	forest, _ := dfs(root, true)
	return forest
}

// // time: O(N)
// // space: O(N)
// func delNodes(root *pkg.TreeNode, toDelete []int) []*pkg.TreeNode {
// 	var inDelMap = make(map[int]bool)
// 	for _, val := range toDelete {
// 		inDelMap[val] = true
// 	}
// 	var dfs func(node *pkg.TreeNode, isRoot bool) []*pkg.TreeNode
// 	dfs = func(node *pkg.TreeNode, isRoot bool) []*pkg.TreeNode {
// 		var forest = make([]*pkg.TreeNode, 0)
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
