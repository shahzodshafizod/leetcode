package trees

import "github.com/shahzodshafizod/leetcode/pkg"

/*
Problem:
Given a binary tree, return the level order traversal of the nodes' values as an array.

Step 1: Verify the constraints
	What do we return if the tree is empty?
		: Return an empty array.
Step 2: Write out some test cases
	- [3, 6, 1, 9, 2, nil, 4, nil, 5, nil, nil, nil, nil, nil, nil, nil, nil, 8, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil]
		: [[3], [6, 1], [9, 2, 4], [5], [8]]
	- [3]
		: [[3]]
	- nil
		: []
	- [1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 5]
		: [[1], [2], [3], [4], [5]]
*/

// https://leetcode.com/problems/binary-tree-level-order-traversal/

// BFS Traversal
func levelOrder(root *pkg.TreeNode) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}

	queue := []*pkg.TreeNode{root}

	for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {
		levelElements := make([]int, 0)

		for count := range queueLen {
			node := queue[count]
			levelElements = append(levelElements, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[queueLen:]

		levels = append(levels, levelElements)
	}

	return levels
}

// // DFS
// func levelOrder(root *pkg.TreeNode) [][]int {
// 	var levels = make([][]int, 0)
// 	var dfs func(root *pkg.TreeNode, level int)
// 	dfs = func(root *pkg.TreeNode, level int) {
// 		if root == nil {
// 			return
// 		}
// 		if len(levels) <= level {
// 			levels = append(levels, []int{root.Val})
// 		} else {
// 			levels[level] = append(levels[level], root.Val)
// 		}
// 		dfs(root.Left, level+1)
// 		dfs(root.Right, level+1)
// 	}
// 	dfs(root, 0)
// 	return levels
// }
