package trees

import (
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

// Approach 1: Store all values, then find min difference
// In-order traversal to get sorted array
// Find minimum difference between adjacent elements
// Time: O(n), Space: O(n)

// Approach 2: In-order Traversal with tracking (Optimal)
// Use BST property: in-order gives sorted sequence
// Track previous node and compute min difference on the fly
// Time: O(n), Space: O(1) excluding recursion stack
func getMinimumDifference(root *pkg.TreeNode) int {
	minDiff := math.MaxInt32

	var prev *pkg.TreeNode

	var inorder func(*pkg.TreeNode)

	inorder = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}

		// Traverse left
		inorder(node.Left)

		// Process current node
		if prev != nil {
			diff := node.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}

		prev = node

		// Traverse right
		inorder(node.Right)
	}

	inorder(root)

	return minDiff
}

// // Alternative: Collect values first
// func getMinimumDifference(root *pkg.TreeNode) int {
// 	values := []int{}

// 	var inorder func(*pkg.TreeNode)

// 	inorder = func(node *pkg.TreeNode) {
// 		if node == nil {
// 			return
// 		}

// 		inorder(node.Left)
// 		values = append(values, node.Val)
// 		inorder(node.Right)
// 	}

// 	inorder(root)

// 	minDiff := math.MaxInt32

// 	for i := 1; i < len(values); i++ {
// 		diff := values[i] - values[i-1]
// 		if diff < minDiff {
// 			minDiff = diff
// 		}
// 	}

// 	return minDiff
// }
