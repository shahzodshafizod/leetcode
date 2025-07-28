package trees

import (
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

/*
Full & Complete Binary Trees

Full Tree: Every node has either two or no (zero) children
	 +--(A)------+
	(B)      +--(C)--+
	        (D)     (E)

Complete Tree: Every level (except the last) is completely full.
Nodes in the last level are filled from the left.
	     +------(A)------+
	 +--(B)--+       +--(C)
	(D)     (E)     (F)

Full & Complete Tree:
	     +------(A)------+
	 +--(B)--+       +--(C)--+
	(D)     (E)     (F)     (G)

Problem:
Given a complete binary tree, count the number of nodes.

Step 1: Verify the constraints
Step 2: Write out some test cases

*/

// https://leetcode.com/problems/count-complete-tree-nodes/

func countNodes(root *pkg.TreeNode) int {
	// find the tree height
	height := 0
	for node := root; node != nil; node = node.Left {
		height++
	}

	if height <= 1 {
		return height
	}

	rightHeight := 0
	for node := root; node != nil; node = node.Right {
		rightHeight++
	}

	// if the binary tree is complete and full
	if height == rightHeight {
		return int(math.Pow(2, float64(height))) - 1
	}

	count := int(math.Pow(2, float64(height-1))) - 1

	optimalRoot, left, right := root, 0, count
	for left < right {
		index := int(math.Ceil(float64(left+right) / 2))

		// find the current node
		currentNode := optimalRoot

		for l, r := left, right; l < r; {
			mid := int(math.Ceil(float64(l+r) / 2))
			if index >= mid {
				currentNode = currentNode.Right
				l = mid
			} else {
				currentNode = currentNode.Left
				r = mid - 1
			}
		}

		if currentNode != nil {
			left = index
			optimalRoot = optimalRoot.Right
		} else {
			right = index - 1
			optimalRoot = optimalRoot.Left
		}
	}

	count += left + 1

	return count
}
