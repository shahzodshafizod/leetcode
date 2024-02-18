package trees

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for (root.Val-p.Val)*(root.Val-q.Val) > 0 {
		if root.Val > p.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root.Val > p.Val && root.Val > q.Val {
// 		return lowestCommonAncestor(root.Left, p, q)
// 	}
// 	if root.Val < p.Val && root.Val < q.Val {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	}
// 	return root
// }
