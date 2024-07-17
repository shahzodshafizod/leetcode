package trees

// https://leetcode.com/problems/delete-nodes-and-return-forest/

// time: O(N)
// space: O(N)
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	var delMap = make(map[int]bool)
	for _, val := range toDelete {
		delMap[val] = true
	}
	var preorder func(node *TreeNode, isRoot bool) []*TreeNode
	preorder = func(node *TreeNode, isRoot bool) []*TreeNode {
		var roots = make([]*TreeNode, 0)
		if isRoot && !delMap[node.Val] {
			roots = append(roots, node)
		}
		if node.Left != nil {
			roots = append(roots, preorder(node.Left, delMap[node.Val])...)
			if delMap[node.Left.Val] {
				node.Left = nil
			}
		}
		if node.Right != nil {
			roots = append(roots, preorder(node.Right, delMap[node.Val])...)
			if delMap[node.Right.Val] {
				node.Right = nil
			}
		}
		return roots
	}
	return preorder(root, true)
}
