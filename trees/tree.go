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

func traversalBFS(root *TreeNode) []int {
	return []int{}
}

func traversalDFS(root *TreeNode) []int {
	return []int{}
}

// Definition for a n-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func makeNAryTree(vals []any) *Node {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	var root = &Node{Val: vals[0].(int), Children: make([]*Node, 0)}
	var parents = []*Node{root}
	var index = 2
	for length := len(parents); length > 0; length = len(parents) {
		for i := 0; i < len(parents); i++ {
			for index < len(vals) && vals[index] != nil {
				var child = &Node{
					Val:      vals[index].(int),
					Children: make([]*Node, 0),
				}
				parents[i].Children = append(parents[i].Children, child)
				parents = append(parents, child)
				index++
			}
			index++
		}
		parents = parents[length:]
	}
	return root
}
