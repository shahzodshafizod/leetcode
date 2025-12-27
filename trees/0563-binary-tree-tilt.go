package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/binary-tree-tilt/

// Approach: Postorder DFS Traversal
// Key insight: For each node, we need the sum of its left and right subtrees.
// Postorder traversal (left, right, root) naturally provides this.
//
// Strategy:
// 1. Use DFS to traverse tree in postorder
// 2. For each node, calculate left subtree sum and right subtree sum
// 3. Tilt = |left_sum - right_sum|
// 4. Accumulate total tilt
// 5. Return current subtree sum to parent
//
// Time Complexity: O(n) where n is number of nodes
// Space Complexity: O(h) where h is height (recursion stack)
func findTilt(root *pkg.TreeNode) int {
	totalTilt := 0

	var dfs func(*pkg.TreeNode) int

	dfs = func(node *pkg.TreeNode) int {
		if node == nil {
			return 0
		}

		// Postorder: process left and right first
		leftSum := dfs(node.Left)
		rightSum := dfs(node.Right)

		// Calculate tilt for current node
		tilt := leftSum - rightSum
		if tilt < 0 {
			tilt = -tilt
		}

		// Accumulate total tilt
		totalTilt += tilt

		// Return sum of current subtree (for parent)
		return leftSum + rightSum + node.Val
	}

	dfs(root)

	return totalTilt
}
