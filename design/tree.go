package design

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(index int, vals []any) *TreeNode {
	if len := len(vals); len == 0 || len <= index || vals[index] == nil {
		return nil
	}
	return &TreeNode{
		Val:   vals[index].(int),
		Left:  MakeTree(2*index+1, vals),
		Right: MakeTree(2*index+2, vals),
	}
}

func MakeTree2(vals ...any) *TreeNode {
	var root *TreeNode = nil
	var queue = make([]*TreeNode, 0)
	if len(vals) > 0 && vals[0] != nil {
		root = &TreeNode{Val: vals[0].(int)}
		queue = append(queue, root)
	}
	var curr *TreeNode = nil
	for idx, n := 1, len(vals); idx < n; idx++ {
		var child *TreeNode = nil
		if vals[idx] != nil {
			child = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, child)
		}
		if curr == nil {
			curr = queue[0]
			queue = queue[1:]
			curr.Left = child
		} else {
			curr.Right = child
			curr = nil
		}
	}
	return root
}

func traversalBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var values = make([]int, 0)
	var queue = NewQueue[*TreeNode]()
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
type TNode struct {
	Val      int
	Children []*TNode
}

func MakeNAryTree(vals []any) *TNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	var root = &TNode{Val: vals[0].(int), Children: make([]*TNode, 0)}
	var parents = []*TNode{root}
	var index = 2
	for length := len(parents); length > 0; length = len(parents) {
		for i := 0; i < len(parents); i++ {
			for index < len(vals) && vals[index] != nil {
				var child = &TNode{
					Val:      vals[index].(int),
					Children: make([]*TNode, 0),
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
