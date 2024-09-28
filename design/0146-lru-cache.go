package design

// https://leetcode.com/problems/lru-cache/

type LRUCache struct {
	head     *dListNode[[2]int]
	tail     *dListNode[[2]int]
	pointers map[int]*dListNode[[2]int]
	length   int
	capacity int
}

func NewLRUCache(capacity int) LRUCache {
	var dummy = &dListNode[[2]int]{}
	return LRUCache{
		head:     dummy,
		tail:     dummy,
		pointers: make(map[int]*dListNode[[2]int]),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	if node := l.pointers[key]; node != nil {
		value := l.unlink(node)
		l.pushback(key, value)
		return value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node := l.pointers[key]; node != nil {
		l.unlink(node)
	}
	if l.length >= l.capacity {
		l.unlink(l.head.next)
	}
	l.pushback(key, value)
}

func (l *LRUCache) unlink(node *dListNode[[2]int]) int {
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev.next = node.next
	if l.tail == node {
		l.tail = l.tail.prev
	}
	delete(l.pointers, node.val[0])
	l.length--
	return node.val[1]
}

func (l *LRUCache) pushback(key int, value int) {
	l.tail.next = NewDListNode([2]int{key, value}, l.tail, nil)
	l.tail = l.tail.next
	l.pointers[key] = l.tail
	l.length++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
