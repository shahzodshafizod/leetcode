package trees

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/create-binary-tree-from-descriptions/

// time: O(2N) = O(N)
// space: O(2N) = O(N)
func createBinaryTree(descriptions [][]int) *design.TreeNode {
	var tree = make(map[int]*design.TreeNode)
	var parent, child, isLeft int
	var children = make(map[int]bool)
	for _, desc := range descriptions {
		parent, child, isLeft = desc[0], desc[1], desc[2]
		children[child] = true
		if tree[parent] == nil {
			tree[parent] = &design.TreeNode{Val: parent}
		}
		if tree[child] == nil {
			tree[child] = &design.TreeNode{Val: child}
		}
		if isLeft == 1 {
			tree[parent].Left = tree[child]
		} else {
			tree[parent].Right = tree[child]
		}
	}
	for _, desc := range descriptions {
		parent = desc[0]
		if !children[parent] {
			return tree[parent]
		}
	}
	return nil
}
