package trees

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/subtree-of-another-tree/

func isSubtree(root *design.TreeNode, subRoot *design.TreeNode) bool {
	var dfs func(root *design.TreeNode, subRoot *design.TreeNode, check bool) bool
	dfs = func(root *design.TreeNode, subRoot *design.TreeNode, check bool) bool {
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

// func isSubtree(root *design.TreeNode, subRoot *design.TreeNode) bool {
// 	var serialize func(node *design.TreeNode) string
// 	serialize = func(node *design.TreeNode) string {
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
