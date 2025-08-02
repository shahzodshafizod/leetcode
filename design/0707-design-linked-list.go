package design

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/design-linked-list/

type MyLinkedList struct {
	head *pkg.DListNode
	tail *pkg.DListNode
	size int
}

func NewMyLinkedList() MyLinkedList {
	ll := MyLinkedList{
		head: &pkg.DListNode{Val: -1},
		tail: &pkg.DListNode{Val: -1},
	}
	ll.head.Next = ll.tail
	ll.tail.Prev = ll.head

	return ll
}

func (m *MyLinkedList) Get(index int) int {
	if node := m.getNodeAtIndex(index); node != nil {
		return node.Val
	}

	return -1
}

func (m *MyLinkedList) AddAtHead(val int) {
	node := &pkg.DListNode{Val: val, Prev: m.head, Next: m.head.Next}
	node.Prev.Next = node
	node.Next.Prev = node
	m.size++
}

func (m *MyLinkedList) AddAtTail(val int) {
	node := &pkg.DListNode{Val: val, Prev: m.tail.Prev, Next: m.tail}
	node.Prev.Next = node
	node.Next.Prev = node
	m.size++
}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index == m.size {
		m.AddAtTail(val)

		return
	}

	next := m.getNodeAtIndex(index)
	if next == nil {
		return
	}

	node := &pkg.DListNode{Val: val, Prev: next.Prev, Next: next}
	node.Prev.Next = node
	node.Next.Prev = node
	m.size++
}

func (m *MyLinkedList) getNodeAtIndex(index int) *pkg.DListNode {
	if index < 0 || index >= m.size {
		return nil
	}

	if index > m.size/2 {
		node := m.tail.Prev
		for i := m.size - 1; i > index; i-- {
			node = node.Prev
		}

		return node
	}

	node := m.head.Next
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	node := m.getNodeAtIndex(index)
	if node == nil {
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	m.size--
}

// type MyLinkedList struct {
// 	head     *pkg.ListNode
// 	pointers []*pkg.ListNode
// 	size     int
// }

// func NewMyLinkedList() MyLinkedList {
// 	return MyLinkedList{pointers: make([]*pkg.ListNode, 0)}
// }

// func (m *MyLinkedList) Get(index int) int {
// 	if !m.validIndex(index) {
// 		return -1
// 	}
// 	return m.pointers[index].Val
// }

// func (m *MyLinkedList) validIndex(index int) bool {
// 	return index >= 0 && index < m.size
// }

// func (m *MyLinkedList) AddAtHead(val int) {
// 	m.head = &pkg.ListNode{Val: val, Next: m.head}
// 	m.pointers = append([]*pkg.ListNode{m.head}, m.pointers...)
// 	m.size++
// }

// func (m *MyLinkedList) AddAtTail(val int) {
// 	newNode := &pkg.ListNode{Val: val, Next: nil}
// 	if m.size > 0 {
// 		m.pointers[m.size-1].Next = newNode
// 	} else {
// 		m.head = newNode
// 	}
// 	m.pointers = append(m.pointers, newNode)
// 	m.size++
// }

// func (m *MyLinkedList) AddAtIndex(index int, val int) {
// 	if index > m.size {
// 		return
// 	}
// 	if index == m.size {
// 		m.AddAtTail(val)
// 		return
// 	}
// 	if index <= 0 {
// 		m.AddAtHead(val)
// 		return
// 	}
// 	before := m.pointers[index-1]
// 	before.Next = &pkg.ListNode{Val: val, Next: before.Next}
// 	m.pointers = append(m.pointers[:index], append([]*pkg.ListNode{before.Next}, m.pointers[index:]...)...)
// 	m.size++
// }

// func (m *MyLinkedList) DeleteAtIndex(index int) {
// 	if !m.validIndex(index) {
// 		return
// 	}
// 	if index == 0 {
// 		m.head = m.head.Next
// 	} else {
// 		before := m.pointers[index-1]
// 		before.Next = before.Next.Next
// 	}
// 	m.pointers = append(m.pointers[:index], m.pointers[index+1:]...)
// 	m.size--
// }

/*
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
