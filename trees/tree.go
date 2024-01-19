package trees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(vals []any) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	vals = vals[1:]
	parents := []*TreeNode{root}
	for len(parents) > 0 && len(vals) > 0 {
		newParents := make([]*TreeNode, 0)
		for _, parent := range parents {
			leftVal := vals[0]
			rightVal := vals[1]
			vals = vals[2:]
			if parent == nil {
				newParents = append(newParents, nil, nil)
				continue
			}
			if leftVal != nil {
				parent.Left = &TreeNode{Val: leftVal.(int)}
			}
			if rightVal != nil {
				parent.Right = &TreeNode{Val: rightVal.(int)}
			}
			newParents = append(newParents, parent.Left)
			newParents = append(newParents, parent.Right)
		}
		parents = newParents
	}
	return root
}
