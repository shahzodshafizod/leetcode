package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/

func minOperations(nums []int, k int) int {
	pq := pkg.NewHeap(nums, func(x, y int) bool { return x < y })
	heap.Init(pq)

	var (
		first  int
		second int
		ok     bool
	)

	count := 0

	for pq.Len() >= 2 && pq.Peak() < k {
		first, ok = heap.Pop(pq).(int)
		_ = ok
		second, ok = heap.Pop(pq).(int)
		_ = ok

		heap.Push(pq, first*2+second)

		count++
	}

	return count
}
