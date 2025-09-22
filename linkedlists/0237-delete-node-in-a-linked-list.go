package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/delete-node-in-a-linked-list/

func deleteNode(node *pkg.ListNode[int]) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
