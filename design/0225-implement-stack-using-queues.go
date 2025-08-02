package design

// https://leetcode.com/problems/implement-stack-using-queues/

type queueNode struct {
	data int
	next *queueNode
}

type MyStack struct {
	first *queueNode
	last  *queueNode
}

func NewMyStack() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(x int) {
	newNode := &queueNode{data: x}
	if m.first == nil {
		m.first, m.last = newNode, newNode
	} else {
		m.last.next = newNode
		m.last = newNode
	}
}

func (m *MyStack) Pop() int {
	var data int
	if m.first == nil {
		return data
	}

	preLast := m.first
	for preLast != m.last && preLast.next != nil && preLast.next != m.last {
		preLast = preLast.next
	}

	if preLast == m.last {
		data = preLast.data
		m.first, m.last = nil, nil
	} else {
		data = preLast.next.data
		m.last = preLast
	}

	return data
}

func (m *MyStack) Top() int {
	var data int
	if m.last != nil {
		data = m.last.data
	}

	return data
}

func (m *MyStack) Empty() bool {
	return m.first == nil
}

// type queueNode struct {
// 	data int
// 	next *queueNode
// }

// type MyStack struct {
// 	inFirst  *queueNode
// 	inLast   *queueNode
// 	outFirst *queueNode
// 	outLast  *queueNode
// }

// func NewMyStack() MyStack {
// 	return MyStack{}
// }

// func (m *MyStack) Push(x int) {
// 	var newNode = &queueNode{data: x}
// 	if m.inFirst == nil {
// 		m.inFirst, m.inLast = newNode, newNode
// 	} else {
// 		m.inLast.next = newNode
// 		m.inLast = newNode
// 	}
// }

// func (m *MyStack) move() {
// 	for m.inFirst != nil {
// 		if m.outFirst == nil {
// 			m.outFirst = m.inFirst
// 		} else {
// 			m.outLast.next = m.inFirst
// 		}
// 		m.outLast = m.inFirst
// 		m.inFirst = m.inFirst.next
// 		m.outLast.next = nil
// 	}
// }

// func (m *MyStack) Pop() int {
// 	var data int
// 	if m.Empty() {
// 		return data
// 	}
// 	var preLast = m.outFirst
// 	for preLast != m.outLast && preLast.next != nil && preLast.next != m.outLast {
// 		preLast = preLast.next
// 	}
// 	if preLast == m.outLast {
// 		data = preLast.data
// 		m.outFirst, m.outLast = nil, nil
// 	} else {
// 		data = preLast.next.data
// 		m.outLast = preLast
// 	}
// 	return data
// }

// func (m *MyStack) Top() int {
// 	var data int
// 	if m.Empty() {
// 		return data
// 	}
// 	if m.outLast != nil {
// 		data = m.outLast.data
// 	}
// 	return data
// }

// func (m *MyStack) Empty() bool {
// 	if m.outFirst == nil {
// 		m.move()
// 	}
// 	return m.outFirst == nil
// }

/*
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
