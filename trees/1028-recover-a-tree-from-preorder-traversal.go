package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

// Approach: Depth-First Search
// Time: O(n)
// Space: O(n)
func recoverFromPreorder(traversal string) *pkg.TreeNode {
	var n = len(traversal)
	var recover func(idx int, depth int) (*pkg.TreeNode, int)
	recover = func(idx int, depth int) (*pkg.TreeNode, int) {
		if idx >= n {
			return nil, idx
		}
		var dashes = 0
		for idx < n && traversal[idx] == '-' {
			idx++
			dashes++
		}
		if dashes != depth {
			return nil, idx - dashes
		}
		var val = 0
		for ; idx < n && traversal[idx] != '-'; idx++ {
			val = val*10 + int(traversal[idx]-'0')
		}
		var node = &pkg.TreeNode{Val: val}
		node.Left, idx = recover(idx, depth+1)
		node.Right, idx = recover(idx, depth+1)
		return node, idx
	}
	var root, _ = recover(0, 0)
	return root
}
