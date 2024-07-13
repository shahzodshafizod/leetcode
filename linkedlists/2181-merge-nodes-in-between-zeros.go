package linkedlists

// https://leetcode.com/problems/merge-nodes-in-between-zeros/

func mergeNodes(head *ListNode) *ListNode {
	var dummy = &ListNode{}
	var tail = dummy
	var sum = 0
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
