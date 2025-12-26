package trees

import (
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/

// Approach: DFS + Sorting
// Track (column, row, value) for each node
// Group by column, sort by row then value
// Time: O(n log n), Space: O(n)
func verticalTraversal(root *pkg.TreeNode) [][]int {
	type nodeInfo struct {
		col, row, val int
	}

	nodes := []nodeInfo{}

	var dfs func(*pkg.TreeNode, int, int)

	dfs = func(node *pkg.TreeNode, row, col int) {
		if node == nil {
			return
		}

		nodes = append(nodes, nodeInfo{col, row, node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}

	dfs(root, 0, 0)

	// Sort by column, then row, then value
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col != nodes[j].col {
			return nodes[i].col < nodes[j].col
		}

		if nodes[i].row != nodes[j].row {
			return nodes[i].row < nodes[j].row
		}

		return nodes[i].val < nodes[j].val
	})

	// Group by column
	result := [][]int{}
	if len(nodes) == 0 {
		return result
	}

	currentCol := nodes[0].col
	currentGroup := []int{nodes[0].val}

	for i := 1; i < len(nodes); i++ {
		if nodes[i].col == currentCol {
			currentGroup = append(currentGroup, nodes[i].val)
		} else {
			result = append(result, currentGroup)
			currentCol = nodes[i].col
			currentGroup = []int{nodes[i].val}
		}
	}

	result = append(result, currentGroup)

	return result
}
