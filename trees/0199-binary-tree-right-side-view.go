package trees

/*
Problem:
Given a binary tree, imagine you're standing to the right of the tree.
Return an array of the values of the nodes you can see ordered from top to bottom.

Step 1: Verify the constraints
Step 2: Write out some test cases
	-				+------(1)--+
		+----------(2)--+		(3)--+
		(4)------+		(5)			(6)
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

func rightSideView(root *TreeNode) []int {
	values := make([]int, 0)
	if root == nil {
		return values
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		count := len(nodes)
		values = append(values, nodes[count-1].Val)
		for count > 0 {
			count--
			node := nodes[0]
			nodes = nodes[1:]
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
	}
	return values
}

// func rightSideViewHelper(root *TreeNode, values *[]int, level int) {
// 	/*
// 		how to use:
// 		values := make([]int, 0)
// 		rightSideViewHelper(root, &values, 1)
// 		return values
// 	*/
// 	if root == nil {
// 		return
// 	}
// 	if len(*values) < level {
// 		*values = append(*values, root.Val)
// 	}
// 	rightSideViewHelper(root.Right, values, level+1)
// 	rightSideViewHelper(root.Left, values, level+1)
// }
