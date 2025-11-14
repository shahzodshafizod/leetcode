package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/intersection-of-two-linked-lists/

func getIntersectionNode(headA, headB *pkg.ListNode) *pkg.ListNode {
	var (
		lenA, lenB   int
		tailA, tailB *pkg.ListNode
	)

	for node := headA; node != nil; node = node.Next {
		lenA++
		tailA = node
	}

	for node := headB; node != nil; node = node.Next {
		lenB++
		tailB = node
	}

	if tailA != tailB {
		return nil
	}

	for ; lenA > lenB; lenA-- {
		headA = headA.Next
	}

	for ; lenB > lenA; lenB-- {
		headB = headB.Next
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
