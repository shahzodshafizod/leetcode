package stacks

type Stack[T any] interface {
	Push(T)
	Pop() T
	Peek() T
	Empty() bool
	Size() int
}

type node[T any] struct {
	val  T
	next *node[T]
}

type stack[T any] struct {
	top  *node[T]
	size int
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(val T) {
	s.top = &node[T]{val: val, next: s.top}
	s.size++
}

func (s *stack[T]) Pop() T {
	var val T
	if !s.Empty() {
		val = s.top.val
		s.top = s.top.next
		s.size--
	}
	return val
}

func (s *stack[T]) Peek() T {
	var val T
	if !s.Empty() {
		val = s.top.val
	}
	return val
}

func (s *stack[T]) Empty() bool {
	return s.top == nil
}

func (s *stack[T]) Size() int {
	return s.size
}
