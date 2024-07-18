package trees

// https://leetcode.com/problems/delete-nodes-and-return-forest/

// time: O(N)
// space: O(N)
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	var inDelMap = make(map[int]bool)
	for _, val := range toDelete {
		inDelMap[val] = true
	}
	var dfs func(node *TreeNode, isRoot bool) ([]*TreeNode, *TreeNode) // (roots, child)
	dfs = func(node *TreeNode, isRoot bool) ([]*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		}
		var forest = make([]*TreeNode, 0)
		if isRoot && !inDelMap[node.Val] {
			forest = append(forest, node)
		}

		var roots []*TreeNode
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
// func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
// 	var inDelMap = make(map[int]bool)
// 	for _, val := range toDelete {
// 		inDelMap[val] = true
// 	}
// 	var dfs func(node *TreeNode, isRoot bool) []*TreeNode
// 	dfs = func(node *TreeNode, isRoot bool) []*TreeNode {
// 		var forest = make([]*TreeNode, 0)
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
