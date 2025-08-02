package design

import (
	"container/list"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/lru-cache/

type LRUCache struct {
	list     list.List
	ptrs     map[int]*list.Element
	capacity int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		ptrs:     make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	if node := l.ptrs[key]; node != nil {
		l.ptrs[key] = l.list.PushFront(l.list.Remove(node))
		return l.ptrs[key].Value.(*pkg.Pair).Val
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node := l.ptrs[key]; node != nil {
		l.list.Remove(node)
	} else {
		l.capacity--
	}

	l.ptrs[key] = l.list.PushFront(&pkg.Pair{Key: key, Val: value})
	if l.capacity < 0 {
		delete(l.ptrs, l.list.Back().Value.(*pkg.Pair).Key)
		l.list.Remove(l.list.Back())

		l.capacity++
	}
}

/*
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
