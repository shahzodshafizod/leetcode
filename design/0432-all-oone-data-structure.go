package design

// https://leetcode.com/problems/all-oone-data-structure/

type allOneNode struct {
	key   string
	count int
	prev  *allOneNode
	next  *allOneNode
}

type AllOne struct {
	head *allOneNode
	tail *allOneNode
	keys map[string]*allOneNode
}

func NewAllOne() AllOne {
	return AllOne{
		keys: make(map[string]*allOneNode),
	}
}

func (a *AllOne) Inc(key string) {
	var current, exists = a.keys[key]
	if !exists {
		current = &allOneNode{key: key, next: a.head}
		if a.head == nil {
			a.head, a.tail = current, current
		} else {
			a.head.prev = current
			a.head = current
		}
		a.keys[key] = current
	}
	current.count++
	for node := current; node != nil && node.next != nil && node.count > node.next.count; node = node.next {
		a.swap(node, node.next)
	}
}

func (a *AllOne) swap(node1 *allOneNode, node2 *allOneNode) {
	a.keys[node1.key], a.keys[node2.key] = a.keys[node2.key], a.keys[node1.key]
	node1.count, node2.count = node2.count, node1.count
	node1.key, node2.key = node2.key, node1.key
}

func (a *AllOne) Dec(key string) {
	var current = a.keys[key]
	current.count--
	if current.count == 0 {
		if current.next != nil {
			current.next.prev = current.prev
		}
		if current.prev != nil {
			current.prev.next = current.next
		}
		if a.head == current {
			a.head = a.head.next
		}
		if a.tail == current {
			a.tail = a.tail.prev
		}
		delete(a.keys, key)
		return
	}
	for node := current; node != nil && node.prev != nil && node.prev.count > node.count; node = node.prev {
		a.swap(node, node.prev)
	}
}

func (a *AllOne) GetMaxKey() string {
	if a.tail == nil {
		return ""
	}
	return a.tail.key
}

func (a *AllOne) GetMinKey() string {
	if a.head == nil {
		return ""
	}
	return a.head.key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
