package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/merge-nodes-in-between-zeros/

func mergeNodes(head *pkg.ListNode[int]) *pkg.ListNode[int] {
	dummy := &pkg.ListNode[int]{}
	tail := dummy
	sum := 0

	for node := head.Next; node != nil; node = node.Next {
		if node.Val != 0 {
			sum += node.Val
		} else {
			node.Val = sum
			sum = 0
			tail.Next = node
			tail = tail.Next
		}
	}

	return dummy.Next
}
