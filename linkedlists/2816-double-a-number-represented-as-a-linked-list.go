package linkedlists

// https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/

func doubleIt(head *ListNode) *ListNode {
	var dummy = &ListNode{Val: 0, Next: head}
	for node := dummy; node.Next != nil; node = node.Next {
		node.Next.Val *= 2
		node.Val += node.Next.Val / 10
		node.Next.Val %= 10
	}
	if dummy.Val != 0 {
		head = dummy
	}
	return head
}

// // time: O(2n) = O(n)
// // space: O(n)
// func doubleIt(head *ListNode) *ListNode {
// 	var stack = make([]*ListNode, 0)
// 	var node *ListNode
// 	for node = head; node != nil; node = node.Next {
// 		stack = append(stack, node)
// 	}
// 	var size = len(stack)
// 	var remainder = 0
// 	for size > 0 {
// 		node = stack[size-1]
// 		size--
// 		remainder += node.Val * 2
// 		node.Val = remainder % 10
// 		remainder /= 10
// 	}
// 	if remainder != 0 {
// 		head = &ListNode{Val: remainder, Next: head}
// 	}
// 	return head
// }

// // time: O(n)
// // space: O(n)
// func doubleIt(head *ListNode) *ListNode {
// 	var double func(node *ListNode) int
// 	double = func(node *ListNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		var remainder int
// 		if node.Next != nil {
// 			remainder = double(node.Next)
// 		}
// 		remainder += node.Val * 2
// 		node.Val = remainder % 10
// 		return remainder / 10
// 	}
// 	var remainder = double(head)
// 	if remainder != 0 {
// 		head = &ListNode{Val: remainder, Next: head}
// 	}
// 	return head
// }
