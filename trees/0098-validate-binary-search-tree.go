package trees

import "math"

/*
Binary Search Trees:
Every left node is lesser than and every right node is greater than the node itself.
		 +------(12)------+
	 +--(8)--+		 +--(18)--+
	(5)		(10)	(14)	(25)

Problem:
Given a binary tree, determine if it is a valid binary search tree.

Step 1: Verify the constraints
	- Can there be duplicate values in the tree?
		: Yes, if you receive duplicate values, the tree is not a valid binary search tree.
Step 2: Write out some test cases:
	-		+------(12)------+				-> true
		+--(8)--+		 +--(18)--+
		(5)		(10)	(14)	(25)

	-		+------(16)------+				-> false
		+--(8)			 +--(22)--+
		(9)				(19)	(25)

	-		+------(13)------+				-> false
		+--(6)			 +--(17)--+
		(2)				(10)	(22)

	-		+------(12)------+				-> true
		+--(7)--+		 +--(18)--+
		(5)		(9)		(16)	(25)
*/

// https://leetcode.com/problems/validate-binary-search-tree/

// time: O(N)
// space: O(N)
func isValidBST(root *TreeNode) bool {
	var prev = math.MinInt
	return inOrderCheck(root, &prev)
}

func inOrderCheck(node *TreeNode, prev *int) bool {
	if node == nil {
		return true
	}

	if !inOrderCheck(node.Left, prev) {
		return false
	}

	if *prev >= node.Val {
		return false
	}
	*prev = node.Val

	return inOrderCheck(node.Right, prev)
}
