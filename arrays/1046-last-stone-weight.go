package arrays

import (
	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/last-stone-weight/

func lastStoneWeight(stones []int) int {
	var maxHeap = pkg.NewPQ(stones, func(x, y int) bool { return x < y })
	maxHeap.Heapify()
	for maxHeap.Len() > 1 {
		y := maxHeap.Pop()
		x := maxHeap.Pop()
		if x != y {
			maxHeap.Push(y - x)
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	}
	return maxHeap.Peek()
}
