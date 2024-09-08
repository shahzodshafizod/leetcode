package linkedlists

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/middle-of-the-linked-list/

func middleNode(head *design.ListNode) *design.ListNode {
	var tortoise, hare = head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	return tortoise
}
