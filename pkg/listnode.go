package pkg

import (
	"reflect"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkedList(vals ...int) *ListNode {
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

// position starts from 0
func GetNode(head *ListNode, position int) *ListNode {
	for node := head; node != nil && position >= 0; {
		if position == 0 {
			return node
		}

		node = node.Next
		position--
	}

	return nil
}

func MakeCycleLinkedList(position int, vals ...int) *ListNode {
	var head, tail *ListNode = nil, nil

	var cyclePoint *ListNode

	for index, val := range vals {
		newNode := &ListNode{Val: val, Next: nil}
		if index == position {
			cyclePoint = newNode
		}

		if head == nil {
			head, tail = newNode, newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}

	if tail != nil {
		tail.Next = cyclePoint
	}

	return head
}

// Definition for a Doubly Linked List Node with Children.
type NAryListNode struct {
	Val   int
	Prev  *NAryListNode
	Next  *NAryListNode
	Child *NAryListNode
}

func MakeNAryLinkedList(vals ...any) *NAryListNode {
	var root, current *NAryListNode

	for _, val := range vals {
		if kind := reflect.TypeOf(val).Kind(); kind != reflect.Int {
			if kind == reflect.Slice && current != nil {
				current.Child = MakeNAryLinkedList(val.([]any)...)
			}

			continue
		}

		newNode := &NAryListNode{Val: val.(int)}
		if root == nil {
			root, current = newNode, newNode
		} else {
			current.Next = newNode
			newNode.Prev = current
			current = newNode
		}
	}

	return root
}
