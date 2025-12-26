package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/find-mode-in-binary-search-tree/

// Approach 1: HashMap/Frequency Counter
// Traverse tree and count frequencies in a map
// Find maximum frequency, then collect all values with that frequency
// Time: O(n), Space: O(n)

// Approach 2: In-order Traversal (Optimal for BST)
// Use BST property: in-order traversal gives sorted values
// Track current value, its count, and max count
// Time: O(n), Space: O(1) excluding recursion stack
func findMode(root *pkg.TreeNode) []int {
	var (
		result []int
		prev   *pkg.TreeNode
	)

	maxCount := 0
	currentCount := 0

	var inorder func(*pkg.TreeNode)

	inorder = func(node *pkg.TreeNode) {
		if node == nil {
			return
		}

		// Traverse left
		inorder(node.Left)

		// Process current node
		if prev == nil || prev.Val != node.Val {
			currentCount = 1
		} else {
			currentCount++
		}

		if currentCount > maxCount {
			maxCount = currentCount
			result = []int{node.Val}
		} else if currentCount == maxCount {
			result = append(result, node.Val)
		}

		prev = node

		// Traverse right
		inorder(node.Right)
	}

	inorder(root)

	return result
}

// // Alternative: Using HashMap
// func findMode(root *pkg.TreeNode) []int {
// 	freq := make(map[int]int)
// 	maxFreq := 0

// 	var traverse func(*pkg.TreeNode)

// 	traverse = func(node *pkg.TreeNode) {
// 		if node == nil {
// 			return
// 		}

// 		freq[node.Val]++
// 		if freq[node.Val] > maxFreq {
// 			maxFreq = freq[node.Val]
// 		}

// 		traverse(node.Left)
// 		traverse(node.Right)
// 	}

// 	traverse(root)

// 	result := []int{}

// 	for val, count := range freq {
// 		if count == maxFreq {
// 			result = append(result, val)
// 		}
// 	}

// 	return result
// }
