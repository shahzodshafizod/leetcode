package trees

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
func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}
	var queue = []*TreeNode{root}
	for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {
		levelElements := make([]int, 0)
		for count := 0; count < queueLen; count++ {
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

// func levelOrderHelper(root *TreeNode, levels *[][]int, level int) {
// 	/*
// 		how to use:
// 		var levels = make([][]int, 0)
// 		levelOrderHelper(root, &levels, 0)
// 		return levels
// 	*/
// 	if root == nil {
// 		return
// 	}
// 	if len(*levels) < level+1 {
// 		*levels = append(*levels, []int{root.Val})
// 	} else {
// 		(*levels)[level] = append((*levels)[level], root.Val)
// 	}
// 	levelOrderHelper(root.Left, levels, level+1)
// 	levelOrderHelper(root.Right, levels, level+1)
// }
