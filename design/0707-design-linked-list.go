package design

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/design-linked-list/

type MyLinkedList struct {
	head *pkg.DListNode[int]
	tail *pkg.DListNode[int]
	size int
}

func NewMyLinkedList() MyLinkedList {
	return MyLinkedList{}
}

func (m *MyLinkedList) Get(index int) int {
	node := m.getNodeAtIndex(index)
	if node == nil {
		return -1
	}
	return node.Val
}

func (m *MyLinkedList) AddAtHead(val int) {
	newNode := pkg.NewDListNode(val, nil, m.head)
	if m.head == nil {
		m.head, m.tail = newNode, newNode
	} else {
		m.head.Prev = newNode
		m.head = newNode
	}
	m.size++
}

func (m *MyLinkedList) AddAtTail(val int) {
	newNode := pkg.NewDListNode(val, m.tail, nil)
	if m.head == nil {
		m.head, m.tail = newNode, newNode
	} else {
		m.tail.Next = newNode
		m.tail = newNode
	}
	m.size++
}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		m.AddAtHead(val)
		return
	}
	if index == m.size {
		m.AddAtTail(val)
		return
	}
	before := m.getNodeAtIndex(index)
	if before == nil {
		return
	}
	var newNode = pkg.NewDListNode(val, before.Prev, before)
	if before.Prev != nil {
		before.Prev.Next = newNode
	}
	before.Prev = newNode
	m.size++
}

func (m *MyLinkedList) getNodeAtIndex(index int) *pkg.DListNode[int] {
	if index < 0 || index >= m.size {
		return nil
	}
	if index > m.size/2 {
		node := m.tail
		for i := m.size - 1; i > index; i-- {
			node = node.Prev
		}
		return node
	}
	node := m.head
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
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node == m.head {
		m.head = m.head.Next
	}
	if node == m.tail {
		m.tail = m.tail.Prev
	}
	m.size--
}

// type MyLinkedList struct {
// 	head     *node[int]
// 	pointers []*node[int]
// 	size     int
// }

// func NewMyLinkedList() MyLinkedList {
// 	return MyLinkedList{pointers: make([]*node[int], 0)}
// }

// func (m *MyLinkedList) Get(index int) int {
// 	if !m.validIndex(index) {
// 		return -1
// 	}
// 	return m.pointers[index].val
// }

// func (m *MyLinkedList) validIndex(index int) bool {
// 	return index >= 0 && index < m.size
// }

// func (m *MyLinkedList) AddAtHead(val int) {
// 	m.head = &node[int]{val: val, next: m.head}
// 	m.pointers = append([]*node[int]{m.head}, m.pointers...)
// 	m.size++
// }

// func (m *MyLinkedList) AddAtTail(val int) {
// 	newNode := &node[int]{val: val, next: nil}
// 	if m.size > 0 {
// 		m.pointers[m.size-1].next = newNode
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
// 	before.next = &node[int]{val: val, next: before.next}
// 	m.pointers = append(m.pointers[:index], append([]*node[int]{before.next}, m.pointers[index:]...)...)
// 	m.size++
// }

// func (m *MyLinkedList) DeleteAtIndex(index int) {
// 	if !m.validIndex(index) {
// 		return
// 	}
// 	if index == 0 {
// 		m.head = m.head.next
// 	} else {
// 		before := m.pointers[index-1]
// 		before.next = before.next.next
// 	}
// 	m.pointers = append(m.pointers[:index], m.pointers[index+1:]...)
// 	m.size--
// }

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
