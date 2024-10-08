package design

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/lru-cache/

type LRUCache struct {
	head     *pkg.DListNode[[2]int]
	tail     *pkg.DListNode[[2]int]
	pointers map[int]*pkg.DListNode[[2]int]
	length   int
	capacity int
}

func NewLRUCache(capacity int) LRUCache {
	var dummy = &pkg.DListNode[[2]int]{}
	return LRUCache{
		head:     dummy,
		tail:     dummy,
		pointers: make(map[int]*pkg.DListNode[[2]int]),
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
		l.unlink(l.head.Next)
	}
	l.pushback(key, value)
}

func (l *LRUCache) unlink(node *pkg.DListNode[[2]int]) int {
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev.Next = node.Next
	if l.tail == node {
		l.tail = l.tail.Prev
	}
	delete(l.pointers, node.Val[0])
	l.length--
	return node.Val[1]
}

func (l *LRUCache) pushback(key int, value int) {
	l.tail.Next = pkg.NewDListNode([2]int{key, value}, l.tail, nil)
	l.tail = l.tail.Next
	l.pointers[key] = l.tail
	l.length++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
