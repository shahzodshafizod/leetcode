package trees

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

func buildTree(preorder []int, inorder []int) *design.TreeNode {
	var indices = make(map[int]int)
	for idx, val := range inorder {
		indices[val] = idx
	}
	var buildTreeRecur func(left, right int) *design.TreeNode
	var idx = 0
	buildTreeRecur = func(left, right int) *design.TreeNode {
		if left > right {
			return nil
		}
		val := preorder[idx]
		idx++
		mid := indices[val]
		return &design.TreeNode{
			Val:   val,
			Left:  buildTreeRecur(left, mid-1),
			Right: buildTreeRecur(mid+1, right),
		}
	}
	return buildTreeRecur(0, len(inorder)-1)
}

// func buildTree(preorder []int, inorder []int) *design.TreeNode {
// 	if len(preorder) == 0 || len(inorder) == 0 {
// 		return nil
// 	}
// 	var val = preorder[0]
// 	preorder = preorder[1:]
// 	var mid = 0
// 	for idx, v := range inorder {
// 		if v == val {
// 			mid = idx
// 			break
// 		}
// 	}
// 	return &design.TreeNode{
// 		Val:   val,
// 		Left:  buildTree(preorder[:mid], inorder[:mid]),
// 		Right: buildTree(preorder[mid:], inorder[mid+1:]),
// 	}
// }
