package linkedlists

// https://leetcode.com/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {
	var minHeap = make([]*ListNode, 0, len(lists))
	var (
		compare = func(i int, j int) bool { return minHeap[i].Val > minHeap[j].Val }
		swap    = func(i int, j int) { minHeap[i], minHeap[j] = minHeap[j], minHeap[i] }
		siftUp  = func() {
			child := len(minHeap) - 1
			parent := (child - 1) / 2
			for parent >= 0 && compare(parent, child) {
				swap(parent, child)
				child = parent
				parent = (child - 1) / 2
			}
		}
		siftDown = func() {
			len := len(minHeap)
			parent := 0
			child := 2*parent + 1
			for child < len {
				if child+1 < len && compare(child, child+1) {
					child++
				}
				if !compare(parent, child) {
					break
				}
				swap(parent, child)
				parent = child
				child = 2*parent + 1
			}
		}
	)
	for _, list := range lists {
		if list != nil {
			minHeap = append(minHeap, list)
			siftUp()
		}
	}
	var list = &ListNode{}
	var tail = list
	for length := len(minHeap); length > 0; length = len(minHeap) {
		swap(0, length-1)
		min := minHeap[length-1]
		minHeap = minHeap[:length-1]
		siftDown()
		if min.Next != nil {
			minHeap = append(minHeap, min.Next)
			siftUp()
		}
		tail.Next = min
		tail = tail.Next
		tail.Next = nil
	}
	return list.Next
}

// func mergeKLists(lists []*ListNode) *ListNode {
// 	var minHeap = design.NewPriorityQueue[*ListNode](func(x, y *ListNode) bool { return x.Val > y.Val })
// 	for _, list := range lists {
// 		if list != nil {
// 			minHeap.Push(list)
// 		}
// 	}
// 	var list = &ListNode{}
// 	var tail = list
// 	for minHeap.Len() > 0 {
// 		min := minHeap.Pop()
// 		if min.Next != nil {
// 			minHeap.Push(min.Next)
// 		}
// 		tail.Next = min
// 		tail = tail.Next
// 		tail.Next = nil
// 	}
// 	return list.Next
// }

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
