package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/even-odd-tree/

func isEvenOddTree(root *pkg.TreeNode) bool {
	queue := []*pkg.TreeNode{root}
	level := 0

	var levelParity, factor int

	for length := len(queue); length > 0; length = len(queue) {
		levelParity = level & 1
		factor = 1

		if levelParity == 1 {
			factor = -1
		}

		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Val&1 == levelParity ||
				i > 0 && (node.Val-queue[i-1].Val)*factor <= 0 {
				return false
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
		queue = queue[length:]
	}

	return true
}
