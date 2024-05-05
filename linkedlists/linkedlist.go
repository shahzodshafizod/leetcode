package linkedlists

import (
	"reflect"
)

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

// position starts from 0
func getNode(head *ListNode, position int) *ListNode {
	for node := head; node != nil && position >= 0; {
		if position == 0 {
			return node
		}
		node = node.Next
		position--
	}
	return nil
}

func makeCycleLinkedList(position int, vals ...int) *ListNode {
	var head, tail *ListNode = nil, nil
	var cyclePoint *ListNode = nil
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

// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func makeDoublyLinkedList(vals ...any) *Node {
	var root, current *Node
	for _, val := range vals {
		if kind := reflect.TypeOf(val).Kind(); kind != reflect.Int {
			if kind == reflect.Slice && current != nil {
				current.Child = makeDoublyLinkedList(val.([]any)...)
			}
			continue
		}
		newNode := &Node{Val: val.(int)}
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
