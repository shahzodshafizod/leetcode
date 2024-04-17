package trees

// https://leetcode.com/problems/smallest-string-starting-from-leaf/

func smallestFromLeaf(root *TreeNode) string {
	var dfs func(node *TreeNode, str string) string
	dfs = func(node *TreeNode, str string) string {
		if node == nil {
			return str
		}
		str = string(byte('a'+node.Val)) + str
		switch {
		case node.Left == nil:
			str = dfs(node.Right, str)
		case node.Right == nil:
			str = dfs(node.Left, str)
		default:
			left := dfs(node.Left, str)
			str = dfs(node.Right, str)
			if left < str {
				str = left
			}
		}
		return str
	}
	return dfs(root, "")
}
