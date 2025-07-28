package linkedlists

import "github.com/shahzodshafizod/leetcode/pkg"

/*
Problem:
Given a linked list and numbers m and n, return it back with only positions m to n in reverse.

Step 1: Verify the constraints
	- Will m and n always be within the bounds of the linked list?
		: Yes, assume 1<=m<=n<=length of the linked list.
	- Can we receive m and n values for the whole linked list?
		: Yes, we can receive m=1 and n=length of the linked list.
Step 2: Write out some test cases
	- (1) -> (2) -> (3) -> (4) -> (5) -> null, m=2, n=4
		: (1) -> (4) -> (3) -> (2) -> (5) -> null
	- (1) -> (2) -> (3) -> (4) -> (5) -> (6) -> null, m=1, n=6
		: (6) -> (5) -> (4) -> (3) -> (2) -> (1) -> null
	- (5) -> null, m=1, n=1
		: (5) -> null
	- null, m=0, n=0
		: null
Steps to reverse linked list:
	1. get current node
	2. store next value
	3. update next value to list so far
	4. store current node of list so far
	5. update current node to stored next value at (2)
	6. increment position
*/

// https://leetcode.com/problems/reverse-linked-list-ii/

func reverseBetween(head *pkg.ListNode, left int, right int) *pkg.ListNode {
	if head == nil {
		return nil
	}

	position := 1
	node := head

	var lastPrev *pkg.ListNode

	for node != nil && position < left {
		lastPrev = node
		node = node.Next
		position++
	}

	var localHead, localTail *pkg.ListNode = nil, node

	for node != nil && position <= right {
		current := node
		node = node.Next
		current.Next = localHead
		localHead = current
		position++
	}

	if lastPrev == nil {
		head = localHead
	} else {
		lastPrev.Next = localHead
	}

	localTail.Next = node

	return head
}
