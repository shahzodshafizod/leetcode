package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

// Approach: Depth-First Search
// Time: O(n)
// Space: O(n)
func recoverFromPreorder(traversal string) *pkg.TreeNode {
	n := len(traversal)

	var restore func(idx int, depth int) (*pkg.TreeNode, int)

	restore = func(idx int, depth int) (*pkg.TreeNode, int) {
		if idx >= n {
			return nil, idx
		}

		dashes := 0

		for idx < n && traversal[idx] == '-' {
			idx++
			dashes++
		}

		if dashes != depth {
			return nil, idx - dashes
		}

		val := 0
		for ; idx < n && traversal[idx] != '-'; idx++ {
			val = val*10 + int(traversal[idx]-'0')
		}

		node := &pkg.TreeNode{Val: val}
		node.Left, idx = restore(idx, depth+1)
		node.Right, idx = restore(idx, depth+1)

		return node, idx
	}
	root, _ := restore(0, 0)

	return root
}
