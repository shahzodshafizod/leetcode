package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/reorder-list/

func reorderList(head *pkg.ListNode[int]) {
	// 1, find middle
	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}

	middle := tortoise

	// 2. reverse the second part
	var tail, next *pkg.ListNode[int]
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
// func reorderList(head *pkg.ListNode) {
// 	var middle, next *pkg.ListNode
// 	var reorder func(*pkg.ListNode, *pkg.ListNode) *pkg.ListNode
// 	reorder = func(left *pkg.ListNode, last *pkg.ListNode) *pkg.ListNode {
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
