package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/merge-in-between-linked-lists/

func mergeInBetween(list1 *pkg.ListNode, a int, b int, list2 *pkg.ListNode) *pkg.ListNode {
	var start *pkg.ListNode
	for node, i := list1, 0; i < a; node, i = node.Next, i+1 {
		start = node
	}

	var end *pkg.ListNode
	for node, i := start, b-a+1; i >= 0; node, i = node.Next, i-1 {
		end = node
	}

	start.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = end.Next

	return list1
}
