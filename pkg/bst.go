package pkg

// Binary Search Tree
type BST interface {
	Search(int) bool
	Insert(int)
	Remove(int)
}

type bst struct {
	root *TreeNode
}

func NewBST() BST {
	return &bst{}
}

func (b *bst) Search(target int) bool {
	return b.search(b.root, target) != nil
}

func (b *bst) search(curr *TreeNode, target int) *TreeNode {
	if curr == nil {
		return nil
	}
	if target < curr.Val {
		return b.search(curr.Left, target)
	}
	if target > curr.Val {
		return b.search(curr.Right, target)
	}
	return curr
}

func (b *bst) Insert(val int) {
	b.root = b.insert(b.root, val)
}

func (b *bst) insert(curr *TreeNode, val int) *TreeNode {
	if curr == nil {
		return &TreeNode{Val: val}
	}
	if val > curr.Val {
		curr.Right = b.insert(curr.Right, val)
	}
	if val < curr.Val {
		curr.Left = b.insert(curr.Left, val)
	}
	return curr
}

func (b *bst) minNodeValue(curr *TreeNode) int {
	var minNode = curr
	for minNode != nil && minNode.Left != nil {
		minNode = minNode.Left
	}
	return minNode.Val
}

func (b *bst) Remove(val int) {
	b.root = b.remove(b.root, val)
}

func (b *bst) remove(curr *TreeNode, val int) *TreeNode {
	if curr == nil {
		return nil
	}
	if val > curr.Val {
		curr.Right = b.remove(curr.Right, val)
	} else if val < curr.Val {
		curr.Left = b.remove(curr.Left, val)
	} else {
		if curr.Left == nil {
			return curr.Right
		}
		if curr.Right == nil {
			return curr.Left
		}
		minNodeVal := b.minNodeValue(curr.Right)
		curr.Val = minNodeVal
		curr.Right = b.remove(curr.Right, minNodeVal)
	}
	return curr
}
