package design

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Peek() T
	Empty() bool
	Size() int
}

type queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(val T) {
	var newNode = &node[T]{val: val}
	if q.head == nil {
		q.head, q.tail = newNode, newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *queue[T]) Dequeue() T {
	var val T
	if !q.Empty() {
		val = q.head.val
		q.head = q.head.next
		q.size--
	}
	return val
}

func (q *queue[T]) Peek() T {
	var val T
	if !q.Empty() {
		val = q.head.val
	}
	return val
}

func (q *queue[T]) Empty() bool {
	return q.head == nil
}

func (q *queue[T]) Size() int {
	return q.size
}
