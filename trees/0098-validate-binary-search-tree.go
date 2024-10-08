package trees

import "github.com/shahzodshafizod/leetcode/pkg"

/*
Binary Search Trees:
Every left node is lesser than and every right node is greater than the node itself.
	     +------(12)------+
	 +--(8)--+       +--(18)--+
	(5)     (10)    (14)     (25)

Problem:
Given a binary tree, determine if it is a valid binary search tree.

Step 1: Verify the constraints
	- Can there be duplicate values in the tree?
		: Yes, if you receive duplicate values, the tree is not a valid binary search tree.
Step 2: Write out some test cases:
	-        +------(12)------+                -> true
	     +--(8)--+       +--(18)--+
	    (5)     (10)    (14)     (25)

	-        +------(16)------+                -> false
	     +--(8)          +--(22)--+
	    (9)             (19)     (25)

	-        +------(13)------+                -> false
	     +--(6)          +--(17)--+
	    (2)             (10)     (22)

	-        +------(12)------+                -> true
	     +--(7)--+       +--(18)--+
	    (5)     (9)     (16)     (25)
*/

// https://leetcode.com/problems/validate-binary-search-tree/

// Approach: Iterative
// N = # of tree nodes
// H = tree height
// time: O(N)
// space: O(H)
func isValidBST(root *pkg.TreeNode) bool {
	var stack = make([]*pkg.TreeNode, 0)
	var prev, node *pkg.TreeNode = nil, root
	for len(stack) != 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev, node = node, node.Right
	}
	return true
}

// // Approach: Recursive
// // N = # of tree nodes
// // H = tree height
// // time: O(N)
// // space: O(H)
// func isValidBST(root *pkg.TreeNode) bool {
// 	var check func(node *pkg.TreeNode, minval int, maxval int) bool
// 	check = func(node *pkg.TreeNode, minval int, maxval int) bool {
// 		if node == nil {
// 			return true
// 		}
// 		if node.Val <= minval || node.Val >= maxval {
// 			return false
// 		}
// 		return check(node.Left, minval, node.Val) &&
// 			check(node.Right, node.Val, maxval)
// 	}
// 	return check(root, math.MinInt, math.MaxInt)
// }
