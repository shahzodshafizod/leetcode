package trees

/*
Binary Tree and Binary Search Trees

Binary Trees - search types:
	- Breadth First Search (BFS)
	- Depth First Search (DFS)
Binary Trees - traversal types:
	- Pre-order traversal (DFS)
	- In-order traversal (DFS)
	- Post-order traversal (DFS)

Problem:
Given a binary tree, find its maximum depth.
Maximum depth is the number of nodes along the longest
path from the root node to the furthest leaf node.

Step 1: Verify the constraints
	- What do we return if the tree is empty?
		: Return 0.
Step 2: Write out some test cases
	-                +---(1)---+
	     +----------(2)       (3)
	    (4)--+
	        (5)--+
	            (6)
	    : 5
	- nil
	    : 0
	- (1)
	    : 1
	- (1)--+
	      (2)--+
	          (3)--+
	              (4)--+
	                  (5)
	    : 5
*/

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
