package linkedlists

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {
	var minHeap = design.NewPriorityQueue[*ListNode](
		make([]*ListNode, 0),
		func(x, y *ListNode) bool { return x.Val > y.Val },
	)
	for _, list := range lists {
		if list != nil {
			minHeap.Push(list)
		}
	}
	var list = &ListNode{}
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

// func mergeKLists(lists []*ListNode) *ListNode {
// 	var minHeap = design.NewHeap[*ListNode](
// 		make([]*ListNode, 0),
// 		func(x, y *ListNode) bool { return x.Val < y.Val },
// 	)
// 	for _, list := range lists {
// 		if list != nil {
// 			heap.Push(minHeap, list)
// 		}
// 	}
// 	var list = &ListNode{}
// 	var tail = list
// 	for minHeap.Len() > 0 {
// 		min := heap.Pop(minHeap).(*ListNode)
// 		if min.Next != nil {
// 			heap.Push(minHeap, min.Next)
// 		}
// 		tail.Next = min
// 		tail = tail.Next
// 		tail.Next = nil
// 	}
// 	return list.Next
// }
