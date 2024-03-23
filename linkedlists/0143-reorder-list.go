package linkedlists

// https://leetcode.com/problems/reorder-list/

func reorderList(head *ListNode) {
	// 1, find middle
	var tortoise, hare = head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	var middle = tortoise

	// 2. reverse the second part
	var tail, next *ListNode
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
// func reorderList(head *ListNode) {
// 	var middle, next *ListNode
// 	var reorder func(*ListNode, *ListNode) *ListNode
// 	reorder = func(left *ListNode, last *ListNode) *ListNode {
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
