package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/remove-linked-list-elements/

// Approach: Iterative
// Time: O(n)
// Space: O(1)
func removeElements(head *pkg.ListNode[int], val int) *pkg.ListNode[int] {
	dummy := &pkg.ListNode[int]{Next: head}

	node := dummy
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return dummy.Next
}

// // Approach: Recursive
// // Time: O(n)
// // Space: O(n)
// func removeElements(head *pkg.ListNode, val int) *pkg.ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	head.Next = removeElements(head.Next, val)
// 	if head.Val == val {
// 		return head.Next
// 	}
// 	return head
// }
