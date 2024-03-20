package linkedlists

// https://leetcode.com/problems/merge-in-between-linked-lists/

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var start *ListNode
	for node, i := list1, 0; i < a; node, i = node.Next, i+1 {
		start = node
	}
	var end *ListNode
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
