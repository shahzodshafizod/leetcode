package design

// https://leetcode.com/problems/all-oone-data-structure/

// Approach #2: Hash Map + Doubly Linked List
type allOneNode struct {
	key   string
	count int
	prev  *allOneNode
	next  *allOneNode
}

func NewAllOneNode(key string, prev *allOneNode, next *allOneNode) *allOneNode {
	return &allOneNode{key: key, prev: prev, next: next}
}

type AllOne struct {
	head *allOneNode
	tail *allOneNode
	keys map[string]*allOneNode
}

func NewAllOne() AllOne {
	head := NewAllOneNode("", nil, nil)
	head.count = 5e4
	tail := NewAllOneNode("", head, nil)
	head.next = tail

	return AllOne{
		keys: make(map[string]*allOneNode),
		head: head,
		tail: tail,
	}
}

func (a *AllOne) Inc(key string) {
	node := a.keys[key]
	if node == nil {
		node = NewAllOneNode(key, a.tail.prev, a.tail)
		node.prev.next = node
		node.next.prev = node
		a.keys[key] = node
	}

	node.count++
	for node.prev.count < node.count {
		a.swap(node.prev, node)
	}
}

func (a *AllOne) Dec(key string) {
	node := a.keys[key]
	node.count--

	if node.count == 0 {
		node.prev.next = node.next
		node.next.prev = node.prev
		a.keys[key] = nil

		return
	}

	for node.count < node.next.count {
		a.swap(node, node.next)
	}
}

func (a *AllOne) GetMaxKey() string {
	if a.head.next == a.tail {
		return ""
	}

	return a.head.next.key
}

func (a *AllOne) GetMinKey() string {
	if a.tail.prev == a.head {
		return ""
	}

	return a.tail.prev.key
}

func (a *AllOne) swap(node1 *allOneNode, node2 *allOneNode) {
	prev := node1.prev
	next := node2.next
	node1.prev, node2.prev = node2, prev
	node1.next, node2.next = next, node1
	prev.next = node2
	next.prev = node1
}

// // Approach #1: Trie
// type allOneNode struct {
// 	count    int
// 	children map[rune]*allOneNode
// }

// func NewAllOneNode() *allOneNode {
// 	return &allOneNode{children: make(map[rune]*allOneNode)}
// }

// type AllOne struct {
// 	root *allOneNode
// }

// func NewAllOne() AllOne {
// 	return AllOne{root: NewAllOneNode()}
// }

// func (a *AllOne) Inc(key string) {
// 	curr := a.root
// 	curr.count++
// 	for _, c := range key {
// 		if curr.children[c] == nil {
// 			curr.children[c] = NewAllOneNode()
// 		}
// 		curr = curr.children[c]
// 		curr.count++
// 	}
// }

// func (a *AllOne) Dec(key string) {
// 	curr := a.root
// 	curr.count--
// 	for _, c := range key {
// 		curr = curr.children[c]
// 		curr.count--
// 	}
// }

// func (a *AllOne) GetMaxKey() string {
// 	var key = ""
// 	curr := a.root
// 	var next *allOneNode
// 	var nextc string
// 	var count int
// 	for curr != nil {
// 		next = nil
// 		nextc = ""
// 		count = 0
// 		for c, child := range curr.children {
// 			if child.count > count {
// 				count = child.count
// 				next = child
// 				nextc = string(c)
// 			}
// 		}
// 		curr = next
// 		key += nextc
// 	}
// 	return key
// }

// func (a *AllOne) GetMinKey() string {
// 	var key = ""
// 	var curr = a.root
// 	var next *allOneNode
// 	var nextc string
// 	var count int
// 	for curr != nil {
// 		next = nil
// 		nextc = ""
// 		count = 5e4
// 		for c, child := range curr.children {
// 			if child.count == 0 {
// 				continue
// 			}
// 			if child.count < count {
// 				count = child.count
// 				next = child
// 				nextc = string(c)
// 			}
// 		}
// 		curr = next
// 		key += nextc
// 	}
// 	return key
// }

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
