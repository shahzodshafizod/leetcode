package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

// Approach #2: Hash Set - One Pass DFS
// Store seen values and check for complement
// N = number of nodes
// Time: O(N) - visit each node once
// Space: O(N) - hash set stores all values
func findTarget(root *pkg.TreeNode, k int) bool {
	seen := make(map[int]bool)

	var dfs func(node *pkg.TreeNode) bool

	dfs = func(node *pkg.TreeNode) bool {
		if node == nil {
			return false
		}
		// Check if complement exists in seen values
		complement := k - node.Val
		if seen[complement] {
			return true
		}
		// Add current value to seen set
		seen[node.Val] = true
		// Check left and right subtrees
		return dfs(node.Left) || dfs(node.Right)
	}

	return dfs(root)
}

// // Approach #1: Brute Force - DFS with Search
// // For each node, search if complement exists in tree
// // N = number of nodes, H = height of tree
// // Time: O(N*H) - visit each node and search for complement
// // Space: O(H) - recursion stack
// func findTarget(root *pkg.TreeNode, k int) bool {
// 	// Search for target value in tree, excluding source node
// 	var search func(node *pkg.TreeNode, target int, source *pkg.TreeNode) bool
// 	search = func(node *pkg.TreeNode, target int, source *pkg.TreeNode) bool {
// 		if node == nil {
// 			return false
// 		}
// 		// Found target but make sure it's not the same node
// 		if node.Val == target && node != source {
// 			return true
// 		}
// 		if target < node.Val {
// 			return search(node.Left, target, source)
// 		}
// 		return search(node.Right, target, source)
// 	}

// 	var dfs func(node *pkg.TreeNode) bool
// 	dfs = func(node *pkg.TreeNode) bool {
// 		if node == nil {
// 			return false
// 		}
// 		// Check if complement exists
// 		complement := k - node.Val
// 		if search(root, complement, node) {
// 			return true
// 		}
// 		// Check left and right subtrees
// 		return dfs(node.Left) || dfs(node.Right)
// 	}

// 	return dfs(root)
// }
