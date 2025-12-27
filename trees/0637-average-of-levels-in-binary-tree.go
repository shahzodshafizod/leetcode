package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/average-of-levels-in-binary-tree/

// Approach: BFS (Breadth-First Search) / Level-Order Traversal
// Time: O(n) - visit each node once
// Space: O(w) - where w is maximum width of tree (queue size)
func averageOfLevels(root *pkg.TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	result := []float64{}
	queue := []*pkg.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		// Process all nodes at current level
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			// Add children for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Calculate average for this level
		result = append(result, float64(levelSum)/float64(levelSize))
	}

	return result
}
