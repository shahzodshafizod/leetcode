package design

/*
Problem:
Implement the class Queue using stacks. The queue methods
you need to implement are enqueue, dequeue, peek, and empty.

Step 1: Verify the constraints
	- Do the queue methods we have to implement need to perform at the same complexity of a real queue?
		: No, but they should be as performant as possible.
Step 2: Write out some test cases
Step 3: Figure out a solution without code

*/

// https://leetcode.com/problems/implement-queue-using-stacks/

type stackNode struct {
	data int
	next *stackNode
}

type MyQueue struct {
	in  *stackNode
	out *stackNode
}

func NewMyQueue() MyQueue {
	return MyQueue{}
}

// enqueue: append a value to the end of the queue
func (q *MyQueue) Push(x int) {
	q.in = &stackNode{data: x, next: q.in}
}

// dequeue: remove the value at the start of the queue
func (q *MyQueue) Pop() int {
	if q.out == nil {
		q.move()
	}

	var result int
	if q.out != nil {
		result = q.out.data
		q.out = q.out.next
	}

	return result
}

func (q *MyQueue) move() {
	for q.in != nil {
		tempNode := q.in
		q.in = q.in.next
		tempNode.next = q.out
		q.out = tempNode
	}
}

// peek: return the value at the start of the queue, not removing it
func (q *MyQueue) Peek() int {
	if q.out == nil {
		q.move()
	}

	var result int
	if q.out != nil {
		result = q.out.data
	}

	return result
}

// empty: return a boolean value of whether the queue is empty or not
func (q *MyQueue) Empty() bool {
	return q.out == nil && q.in == nil
}
