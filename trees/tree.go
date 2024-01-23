package trees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(index int, vals []any) *TreeNode {
	if len := len(vals); len == 0 || len <= index || vals[index] == nil {
		return nil
	}
	return &TreeNode{
		Val:   vals[index].(int),
		Left:  makeTree(2*index+1, vals),
		Right: makeTree(2*index+2, vals),
	}
}
