package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/

// Approach: Recursive DFS
// Recursively find the maximum depth among all children and add 1 for current node.
// Time: O(n) - visit each node once
// Space: O(h) - recursion stack where h is height
func maxDepthNary(root *pkg.NTreeNode) int {
	if root == nil {
		return 0
	}

	maxChildDepth := 0

	for _, child := range root.Children {
		childDepth := maxDepthNary(child)
		if childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}

	return 1 + maxChildDepth
}
