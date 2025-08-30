package trees

import "github.com/shahzodshafizod/leetcode/pkg"

/*
Problem:
Given a binary tree, imagine you're standing to the right of the tree.
Return an array of the values of the nodes you can see ordered from top to bottom.

Step 1: Verify the constraints
Step 2: Write out some test cases
	-                +------(1)--+
	     +----------(2)--+      (3)--+
	    (4)------+      (5)         (6)
	         +--(7)
	        (8)
		: [1, 3, 6, 7, 8]
	- (1)
		: [1]
Depth-First Search:
	PreOrder:  NLR --(switch order)-> NRL (the solution for DFS)
	InOrder:   LNR --(switch order)-> RNL
	PostOrder: LRN --(switch order)-> RLN
*/

// https://leetcode.com/problems/binary-tree-right-side-view/

// BFS Traversal
func rightSideView(root *pkg.TreeNode) []int {
	values := make([]int, 0)
	if root == nil {
		return values
	}

	queue := []*pkg.TreeNode{root}
	for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {
		values = append(values, queue[queueLen-1].Val)

		for count := range queueLen {
			node := queue[count]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[queueLen:]
	}

	return values
}

// func rightSideView(root *pkg.TreeNode) []int {
// 	values := make([]int, 0)
// 	var dfs func(root *pkg.TreeNode, level int)
// 	dfs = func(root *pkg.TreeNode, level int) {
// 		if root == nil {
// 			return
// 		}
// 		if len(values) < level {
// 			values = append(values, root.Val)
// 		}
// 		dfs(root.Right, level+1)
// 		dfs(root.Left, level+1)
// 	}
// 	dfs(root, 1)
// 	return values
// }
