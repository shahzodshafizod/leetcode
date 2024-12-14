package design

import (
	"container/list"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/lfu-cache/

type LFUCache struct {
	capacity int
	nodes    map[int]*list.Element
	counts   map[int]*list.List
	minCnt   int
}

func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodes:    make(map[int]*list.Element),
		counts:   map[int]*list.List{1: {}},
		minCnt:   1,
	}
}

func (l *LFUCache) Get(key int) int {
	if l.nodes[key] == nil {
		return -1
	}
	var node = l.nodes[key]
	var pair = node.Value.(*pkg.Pair)
	l.counts[pair.Cnt].Remove(node)
	if l.counts[pair.Cnt].Len() == 0 && l.minCnt == pair.Cnt {
		l.minCnt++
	}
	pair.Cnt++
	if l.counts[pair.Cnt] == nil {
		l.counts[pair.Cnt] = &list.List{}
	}
	l.nodes[key] = l.counts[pair.Cnt].PushFront(pair)
	return pair.Val
}

func (l *LFUCache) Put(key int, value int) {
	if l.Get(key) != -1 {
		var pair = l.nodes[key].Value.(*pkg.Pair)
		pair.Val = value
		l.nodes[key].Value = pair
		return
	}
	if l.capacity == 0 {
		var node = l.counts[l.minCnt].Back()
		var pair = node.Value.(*pkg.Pair)
		delete(l.nodes, pair.Key)
		l.counts[l.minCnt].Remove(node)
		l.capacity++
	}
	l.nodes[key] = l.counts[1].PushFront(&pkg.Pair{Key: key, Val: value, Cnt: 1})
	l.minCnt = 1
	l.capacity--
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
