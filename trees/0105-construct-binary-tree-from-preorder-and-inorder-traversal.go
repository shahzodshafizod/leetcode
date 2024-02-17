package trees

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

func buildTree(preorder []int, inorder []int) *TreeNode {
	var indices = make(map[int]int)
	for idx, val := range inorder {
		indices[val] = idx
	}
	var buildTreeRecur func(left, right int) *TreeNode
	var idx = 0
	buildTreeRecur = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		val := preorder[idx]
		idx++
		mid := indices[val]
		return &TreeNode{
			Val:   val,
			Left:  buildTreeRecur(left, mid-1),
			Right: buildTreeRecur(mid+1, right),
		}
	}
	return buildTreeRecur(0, len(inorder)-1)
}

// func buildTree(preorder []int, inorder []int) *TreeNode {
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
// 	return &TreeNode{
// 		Val:   val,
// 		Left:  buildTree(preorder[:mid], inorder[:mid]),
// 		Right: buildTree(preorder[mid:], inorder[mid+1:]),
// 	}
// }
