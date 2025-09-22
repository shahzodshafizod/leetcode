package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates(head *pkg.ListNode[int]) *pkg.ListNode[int] {
	dummy := &pkg.ListNode[int]{Next: head, Val: -10000}
	for node := dummy; node.Next != nil; {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return dummy.Next
}
