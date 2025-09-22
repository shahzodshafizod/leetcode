package linkedlists

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/merge-k-sorted-lists/

func mergeKLists(lists []*pkg.ListNode[int]) *pkg.ListNode[int] {
	minHeap := pkg.NewHeap(
		make([]*pkg.ListNode[int], 0),
		func(x, y *pkg.ListNode[int]) bool { return x.Val < y.Val },
	)

	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	list := &pkg.ListNode[int]{}
	tail := list

	for minHeap.Len() > 0 {
		minimum, ok := heap.Pop(minHeap).(*pkg.ListNode[int])
		_ = ok

		if minimum.Next != nil {
			heap.Push(minHeap, minimum.Next)
		}

		tail.Next = minimum
		tail = tail.Next
		tail.Next = nil
	}

	return list.Next
}

// func mergeKLists(lists []*pkg.ListNode) *pkg.ListNode {
// 	var minHeap = design.NewHeap[*pkg.ListNode](
// 		make([]*pkg.ListNode, 0),
// 		func(x, y *pkg.ListNode) bool { return x.Val < y.Val },
// 	)
// 	for _, list := range lists {
// 		if list != nil {
// 			heap.Push(minHeap, list)
// 		}
// 	}
// 	var list = &pkg.ListNode{}
// 	var tail = list
// 	for minHeap.Len() > 0 {
// 		min := heap.Pop(minHeap).(*pkg.ListNode)
// 		if min.Next != nil {
// 			heap.Push(minHeap, min.Next)
// 		}
// 		tail.Next = min
// 		tail = tail.Next
// 		tail.Next = nil
// 	}
// 	return list.Next
// }
