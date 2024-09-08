package linkedlists

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *design.ListNode, n int) *design.ListNode {
	var before = &design.ListNode{Next: head}
	var current = before
	for n > 0 {
		n--
		current = current.Next
	}
	for current.Next != nil {
		current = current.Next
		before = before.Next
	}
	current = before.Next
	before.Next = current.Next
	if head == current {
		head = head.Next
	}
	return head
}
