package design

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/design-a-number-container-system/

type NumberContainers struct {
	idx2num map[int]int
	num2idx map[int]*pkg.Heap[int]
}

func NewNumberContainers() NumberContainers {
	return NumberContainers{
		idx2num: make(map[int]int),
		num2idx: make(map[int]*pkg.Heap[int]),
	}
}

func (n *NumberContainers) Change(index int, number int) {
	n.idx2num[index] = number
	if n.num2idx[number] == nil {
		n.num2idx[number] = pkg.NewHeap(make([]int, 0), func(x int, y int) bool { return x < y })
	}

	heap.Push(n.num2idx[number], index)
}

func (n *NumberContainers) Find(number int) int {
	for n.num2idx[number] != nil && n.num2idx[number].Len() > 0 && n.idx2num[n.num2idx[number].Peak()] != number {
		heap.Pop(n.num2idx[number])
	}

	if n.num2idx[number] != nil && n.num2idx[number].Len() > 0 {
		return n.num2idx[number].Peak()
	}

	return -1
}

/*
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
