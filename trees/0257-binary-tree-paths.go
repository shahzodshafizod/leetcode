package trees

import (
	"strconv"
	"strings"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/binary-tree-paths/

// Approach: Depth-First Search + Backtracking
// Time: O(n), n=# of nodes
// Space: O(h), h=height of tree
func binaryTreePaths(root *pkg.TreeNode) []string {
	var dfs func(node *pkg.TreeNode, path []string) []string

	dfs = func(node *pkg.TreeNode, path []string) []string {
		if node == nil {
			return nil
		}

		path = append(path, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			return []string{strings.Join(path, "->")}
		}

		paths := dfs(node.Left, path)
		paths = append(paths, dfs(node.Right, path)...)

		return paths
	}

	return dfs(root, []string{})
}
