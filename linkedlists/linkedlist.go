package linkedlists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func makeLinkedList(vals ...int) *ListNode {
	var head, tail *ListNode = nil, nil
	for _, val := range vals {
		newNode := &ListNode{Val: val, Next: nil}
		if head == nil {
			head, tail = newNode, newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}
	return head
}
