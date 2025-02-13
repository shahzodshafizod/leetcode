package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/

func minOperations(nums []int, k int) int {
	var pq = pkg.NewHeap(nums, func(x, y int) bool { return x < y })
	heap.Init(pq)
	var first, second int
	var count = 0
	for pq.Len() >= 2 && pq.Peek() < k {
		first = heap.Pop(pq).(int)
		second = heap.Pop(pq).(int)
		heap.Push(pq, first*2+second)
		count++
	}
	return count
}
