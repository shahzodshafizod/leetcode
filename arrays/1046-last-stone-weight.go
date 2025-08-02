package arrays

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/last-stone-weight/

func lastStoneWeight(stones []int) int {
	maxHeap := pkg.NewHeap(stones, func(x, y int) bool { return x > y })
	heap.Init(maxHeap)

	for maxHeap.Len() > 1 {
		y, ok := heap.Pop(maxHeap).(int)
		_ = ok
		x, ok := heap.Pop(maxHeap).(int)
		_ = ok

		if x != y {
			heap.Push(maxHeap, y-x)
		}
	}

	if maxHeap.Len() == 0 {
		return 0
	}

	return maxHeap.Peak()
}
