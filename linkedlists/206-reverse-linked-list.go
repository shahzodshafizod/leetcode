package linkedlists

/*
Problem:
Given a linked list, return it in reverse.

Step 1: Verify the constraints
	- What do we return if we get null or a single node?
		: Just return null and the node back respectively.
Step 2: Write out some test cases
	- (1) -> (2) -> (3) -> (4) -> (5) -> null
		: (5) -> (4) -> (3) -> (2) -> (1) -> null
	- (3) -> null
		: (3) -> null
	- null
		: null
*/

// https://leetcode.com/problems/reverse-linked-list/

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for i := head; i != nil; {
		prev := newHead
		newHead = i
		i = i.Next
		newHead.Next = prev
	}
	return newHead
}

// func reverseList(head *ListNode) *ListNode {
// 	var newHead *ListNode = nil
// 	for i := head; i != nil; i = i.Next {
// 		newHead = &ListNode{Val: i.Val, Next: newHead}
// 	}
// 	return newHead
// }
