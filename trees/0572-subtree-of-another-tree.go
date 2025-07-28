package trees

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/subtree-of-another-tree/

func isSubtree(root *pkg.TreeNode, subRoot *pkg.TreeNode) bool {
	var dfs func(root *pkg.TreeNode, subRoot *pkg.TreeNode, check bool) bool

	dfs = func(root *pkg.TreeNode, subRoot *pkg.TreeNode, check bool) bool {
		if root == nil || subRoot == nil {
			return root == subRoot
		}

		if check {
			return root.Val == subRoot.Val &&
				dfs(root.Left, subRoot.Left, true) &&
				dfs(root.Right, subRoot.Right, true)
		}

		return dfs(root, subRoot, true) ||
			dfs(root.Left, subRoot, false) ||
			dfs(root.Right, subRoot, false)
	}

	return dfs(root, subRoot, false)
}

// func isSubtree(root *pkg.TreeNode, subRoot *pkg.TreeNode) bool {
// 	var serialize func(node *pkg.TreeNode) string
// 	serialize = func(node *pkg.TreeNode) string {
// 		if node == nil {
// 			return ";"
// 		}
// 		return fmt.Sprintf("S%dL%sR%sE",
// 			node.Val,
// 			serialize(node.Left),
// 			serialize(node.Right),
// 		)
// 	}
// 	return strings.Contains(serialize(root), serialize(subRoot))
// }
