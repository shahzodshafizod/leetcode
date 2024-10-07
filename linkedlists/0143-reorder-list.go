package linkedlists

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/reorder-list/

func reorderList(head *design.ListNode) {
	// 1, find middle
	var tortoise, hare = head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	var middle = tortoise

	// 2. reverse the second part
	var tail, next *design.ListNode
	for tortoise != nil {
		next = tortoise.Next
		tortoise.Next = tail
		tail = tortoise
		tortoise = next
	}

	// 3. reorder
	for head != middle {
		next = head.Next
		head.Next, head = tail, tail
		tail = next
	}
}

// // recursion
// func reorderList(head *design.ListNode) {
// 	var middle, next *design.ListNode
// 	var reorder func(*design.ListNode, *design.ListNode) *design.ListNode
// 	reorder = func(left *design.ListNode, last *design.ListNode) *design.ListNode {
// 		if last == nil || last.Next == nil {
// 			middle = left
// 			if last != nil {
// 				left = left.Next
// 			}
// 			return left
// 		}
// 		right := reorder(left.Next, last.Next.Next)
// 		next = right.Next
// 		right.Next = left.Next
// 		left.Next = right
// 		return next
// 	}
// 	reorder(head, head)
// 	middle.Next = nil
// }
