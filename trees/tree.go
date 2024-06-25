package trees

import (
	"github.com/shahzodshafizod/alkhwarizmi/design"
)

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
	if root == nil {
		return []int{}
	}
	var values = make([]int, 0)
	var queue = design.NewQueue[*TreeNode]()
	queue.Enqueue(root)
	for queue.Size() > 0 {
		current := queue.Dequeue()
		values = append(values, current.Val)
		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	return values
}

func traversalDFSInOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var values = make([]int, 0)
	if root.Left != nil {
		values = append(values, traversalDFSInOrder(root.Left)...)
	}
	values = append(values, root.Val)
	if root.Right != nil {
		values = append(values, traversalDFSInOrder(root.Right)...)
	}
	return values
}

func traversalMorris(root *TreeNode) []int {
	var values = make([]int, 0)
	var prev *TreeNode
	var curr = root
	for curr != nil {
		if curr.Left == nil {
			values = append(values, curr.Val)
			curr = curr.Right
		} else {
			prev = curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}
			if prev.Right == curr {
				prev.Right = nil
				values = append(values, curr.Val)
				curr = curr.Right
			} else {
				prev.Right = curr
				curr = curr.Left
			}
		}
	}
	return values
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
