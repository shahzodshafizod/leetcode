package trees

import (
	"strings"

	"github.com/shahzodshafizod/leetcode/design"
)

// LCA - Lowest Common Ancestor

// https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/

// time: O(N)
// space: O(N)
func getDirections(root *design.TreeNode, startValue int, destValue int) string {
	var dfs func(node *design.TreeNode, target int, level int, direction byte) []byte
	dfs = func(node *design.TreeNode, target int, level int, direction byte) []byte {
		if node == nil {
			return nil
		}
		if node.Val == target {
			var path = make([]byte, level+1)
			path[level] = direction
			return path
		}
		var path = dfs(node.Left, target, level+1, 'L')
		if len(path) > 0 {
			path[level] = direction
			return path
		}
		path = dfs(node.Right, target, level+1, 'R')
		if len(path) > 0 {
			path[level] = direction
			return path
		}
		return nil
	}
	var startPath = dfs(root, startValue, 0, '0')
	var destPath = dfs(root, destValue, 0, '0')
	var idx, minlen = 0, min(len(startPath), len(destPath))
	for idx < minlen && startPath[idx] == destPath[idx] {
		idx++
	}
	var directions = strings.Repeat("U", len(startPath)-idx) + string(destPath[idx:])
	return directions
}
