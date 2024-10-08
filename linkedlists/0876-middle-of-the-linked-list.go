package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/middle-of-the-linked-list/

func middleNode(head *pkg.ListNode) *pkg.ListNode {
	var tortoise, hare = head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	return tortoise
}
