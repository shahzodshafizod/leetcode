package design

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/

type FindElements struct {
	values map[int]struct{}
}

func NewFindElements(root *pkg.TreeNode) FindElements {
	var values = make(map[int]struct{})
	root.Val = 0
	var queue = []*pkg.TreeNode{root}
	for len(queue) > 0 {
		var next = make([]*pkg.TreeNode, 0)
		for _, node := range queue {
			values[node.Val] = struct{}{}
			if node.Left != nil {
				node.Left.Val = node.Val*2 + 1
				next = append(next, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*2 + 2
				next = append(next, node.Right)
			}
		}
		queue = next
	}
	return FindElements{values}
}

func (f *FindElements) Find(target int) bool {
	var _, found = f.values[target]
	return found
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
