package design

// Binary Search Tree
type BST interface {
	Search(int) bool
	Insert(int)
	Remove(int)
}

type bst struct {
	root *treeNode
}

func NewBST() BST {
	return &bst{}
}

func (b *bst) Search(target int) bool {
	return b.search(b.root, target) != nil
}

func (b *bst) search(curr *treeNode, target int) *treeNode {
	if curr == nil {
		return nil
	}
	if target < curr.val {
		return b.search(curr.left, target)
	}
	if target > curr.val {
		return b.search(curr.right, target)
	}
	return curr
}

func (b *bst) Insert(val int) {
	b.root = b.insert(b.root, val)
}

func (b *bst) insert(curr *treeNode, val int) *treeNode {
	if curr == nil {
		return &treeNode{val: val}
	}
	if val > curr.val {
		curr.right = b.insert(curr.right, val)
	}
	if val < curr.val {
		curr.left = b.insert(curr.left, val)
	}
	return curr
}

func (b *bst) minNodeValue(curr *treeNode) int {
	var minNode = curr
	for minNode != nil && minNode.left != nil {
		minNode = minNode.left
	}
	return minNode.val
}

func (b *bst) Remove(val int) {
	b.root = b.remove(b.root, val)
}

func (b *bst) remove(curr *treeNode, val int) *treeNode {
	if curr == nil {
		return nil
	}
	if val > curr.val {
		curr.right = b.remove(curr.right, val)
	} else if val < curr.val {
		curr.left = b.remove(curr.left, val)
	} else {
		if curr.left == nil {
			return curr.right
		}
		if curr.right == nil {
			return curr.left
		}
		minNodeVal := b.minNodeValue(curr.right)
		curr.val = minNodeVal
		curr.right = b.remove(curr.right, minNodeVal)
	}
	return curr
}
