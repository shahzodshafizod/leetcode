package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/n-ary-tree-level-order-traversal/

func levelOrder429(root *pkg.NTreeNode) [][]int {
	var values = make([][]int, 0)
	var queue = make([]*pkg.NTreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for length := len(queue); length > 0; length = len(queue) {
		var levelValues = make([]int, 0)
		for i := 0; i < length; i++ {
			levelValues = append(levelValues, queue[i].Val)
			queue = append(queue, queue[i].Children...)
		}
		queue = queue[length:]
		values = append(values, levelValues)
	}
	return values
}
