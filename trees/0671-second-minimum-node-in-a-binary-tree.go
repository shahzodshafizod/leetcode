package trees

import (
	"math"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/

// Approach: Optimized DFS - Use tree property
// Key insight: root.val is always the minimum (by tree property)
// We only need to find the smallest value > root.val
// Prune branches where node.val == root.val (can't have second min there)
// N = number of nodes
// Time: O(N) - worst case visit all nodes, but often prunes early
// Space: O(H) - recursion stack depth H (tree height)
func findSecondMinimumValue(root *pkg.TreeNode) int {
	if root == nil {
		return -1
	}

	minVal := root.Val
	secondMin := math.MaxInt

	var dfs func(node *pkg.TreeNode)

	dfs = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}

		// If current value is greater than min but less than secondMin
		if minVal < node.Val && node.Val < secondMin {
			secondMin = node.Val
		}

		// Only explore children if they might contain second minimum
		// If node.val == minVal, children might have larger values
		// If node.val > minVal, we already updated secondMin
		if node.Val == minVal {
			dfs(node.Left)
			dfs(node.Right)
		}
	}

	dfs(root)

	if secondMin == math.MaxInt {
		return -1
	}

	return secondMin
}
