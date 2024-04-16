package trees

// https://leetcode.com/problems/add-one-row-to-tree/

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	var dfs func(node *TreeNode, val int, depth int)
	dfs = func(node *TreeNode, val int, depth int) {
		if depth == 1 {
			node.Left = &TreeNode{Val: val, Left: node.Left}
			node.Right = &TreeNode{Val: val, Right: node.Right}
			return
		}
		if node.Left != nil {
			dfs(node.Left, val, depth-1)
		}
		if node.Right != nil {
			dfs(node.Right, val, depth-1)
		}
	}
	if depth == 1 {
		root = &TreeNode{Val: val, Left: root}
	} else {
		dfs(root, val, depth-1)
	}
	return root
}
