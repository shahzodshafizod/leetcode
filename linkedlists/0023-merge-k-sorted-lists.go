package linkedlists

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/merge-k-sorted-lists/

func mergeKLists(lists []*design.ListNode) *design.ListNode {
	var minHeap = design.NewPQ(
		make([]*design.ListNode, 0),
		func(x, y *design.ListNode) bool { return x.Val > y.Val },
	)
	for _, list := range lists {
		if list != nil {
			minHeap.Push(list)
		}
	}
	var list = &design.ListNode{}
	var tail = list
	for minHeap.Len() > 0 {
		min := minHeap.Pop()
		if min.Next != nil {
			minHeap.Push(min.Next)
		}
		tail.Next = min
		tail = tail.Next
		tail.Next = nil
	}
	return list.Next
}

// func mergeKLists(lists []*design.ListNode) *design.ListNode {
// 	var minHeap = design.NewHeap[*design.ListNode](
// 		make([]*design.ListNode, 0),
// 		func(x, y *design.ListNode) bool { return x.Val < y.Val },
// 	)
// 	for _, list := range lists {
// 		if list != nil {
// 			heap.Push(minHeap, list)
// 		}
// 	}
// 	var list = &design.ListNode{}
// 	var tail = list
// 	for minHeap.Len() > 0 {
// 		min := heap.Pop(minHeap).(*design.ListNode)
// 		if min.Next != nil {
// 			heap.Push(minHeap, min.Next)
// 		}
// 		tail.Next = min
// 		tail = tail.Next
// 		tail.Next = nil
// 	}
// 	return list.Next
// }
