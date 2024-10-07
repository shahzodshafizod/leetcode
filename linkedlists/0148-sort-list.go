package linkedlists

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/sort-list/

// Approach#5: Merge Sort
// time: O(n log n)
// space: O(log n) - for recursion stack OR w/o: O(1)
func sortList(head *design.ListNode) *design.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var tortoise, hare = head, head.Next
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	var secondPart = tortoise.Next
	tortoise.Next = nil
	return mergeTwoLists(
		sortList(head),
		sortList(secondPart),
	)
}

// // Approach#4: Bucket (Count) Sort
// // time: O(3n) = O(n)
// // space: O(n)
// func sortList(head *design.ListNode) *design.ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	maximum, minimum := head.Val, head.Val
// 	for node := head; node != nil; node = node.Next {
// 		maximum = max(maximum, node.Val)
// 		minimum = min(minimum, node.Val)
// 	}
// 	counts := make([]uint, maximum-minimum+1)
// 	for node := head; node != nil; node = node.Next {
// 		counts[node.Val-minimum]++
// 	}
// 	var node = head
// 	for val, count := range counts {
// 		for count > 0 {
// 			count--
// 			node.Val = val + minimum
// 			node = node.Next
// 		}
// 	}
// 	return head
// }

// // Approach#3: Sorting
// // time: O(n log n)
// // space: O(n)
// func sortList(head *design.ListNode) *design.ListNode {
// 	var vals = make([]int, 0)
// 	for i := head; i != nil; i = i.Next {
// 		vals = append(vals, i.Val)
// 	}
// 	sort.Ints(vals)
// 	var idx = 0
// 	for node := head; node != nil; node = node.Next {
// 		node.Val = vals[idx]
// 		idx++
// 	}
// 	return head
// }

// // Approach#2: Priority Queue
// // time: O(n log n)
// // space: O(n)
// func sortList(head *design.ListNode) *design.ListNode {
// 	var pq = design.NewPQ[int](make([]int, 0), func(x, y int) bool { return x > y })
// 	for node := head; node != nil; node = node.Next { // O(n)
// 		pq.Push(node.Val) // O(log n)
// 	}
// 	for node := head; node != nil; node = node.Next { // O(n)
// 		node.Val = pq.Pop()
// 	}
// 	return head
// }

// // Approach#1: Selection Sort
// // time: O(n ^ 2) // Time Limit Exceeded
// // space: O(1)
// func sortList(head *design.ListNode) *design.ListNode {
// 	for i := head; i != nil; i = i.Next {
// 		for j := i.Next; j != nil; j = j.Next {
// 			if i.Val > j.Val {
// 				i.Val, j.Val = j.Val, i.Val
// 			}
// 		}
// 	}
// 	return head
// }
