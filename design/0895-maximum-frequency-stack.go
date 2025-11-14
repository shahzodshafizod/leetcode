package design

import (
	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/maximum-frequency-stack/

// Heap (Priority Queue) will not work,
// because it cannot keep track of order:
// "If there is a tie for the most frequent element,
// the element closest to the stack's top is removed and returned."

type FreqStack struct {
	counts  map[int]int
	buckets map[int]*pkg.ListNode
	maxCnt  int
}

func NewFreqStack() FreqStack {
	return FreqStack{
		counts:  make(map[int]int),
		buckets: make(map[int]*pkg.ListNode),
	}
}

func (f *FreqStack) Push(val int) {
	f.counts[val]++
	cnt := f.counts[val]
	f.buckets[cnt] = &pkg.ListNode{Val: val, Next: f.buckets[cnt]}
	f.maxCnt = max(f.maxCnt, cnt)
}

func (f *FreqStack) Pop() int {
	val := f.buckets[f.maxCnt].Val
	f.counts[val]--

	f.buckets[f.maxCnt] = f.buckets[f.maxCnt].Next
	if f.buckets[f.maxCnt] == nil {
		f.maxCnt--
	}

	return val
}

/*
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
