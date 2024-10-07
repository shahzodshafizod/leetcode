package linkedlists

import "github.com/shahzodshafizod/leetcode/design"

// https://leetcode.com/problems/delete-node-in-a-linked-list/

func deleteNode(node *design.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
