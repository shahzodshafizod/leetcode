package trees

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

func recoverFromPreorder(traversal string) *design.TreeNode {
	var idx = 0
	return recoverFromPreorderRecur(traversal, &idx, len(traversal), 0)
}

func recoverFromPreorderRecur(traversal string, idx *int, n int, level int) *design.TreeNode {
	if *idx >= n {
		return nil
	}
	var i = *idx
	for dashes := 0; dashes < level; dashes++ {
		if traversal[i] != '-' {
			return nil
		}
		i++
	}
	*idx = i
	var val = 0
	for *idx < n && traversal[*idx] != '-' {
		val = val*10 + int(traversal[*idx]-'0')
		(*idx)++
	}
	var node = &design.TreeNode{Val: val}
	node.Left = recoverFromPreorderRecur(traversal, idx, n, level+1)
	node.Right = recoverFromPreorderRecur(traversal, idx, n, level+1)
	return node
}
