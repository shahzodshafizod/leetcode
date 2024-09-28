package design

type listNode[T any] struct {
	val  T
	next *listNode[T]
}

type dListNode[T any] struct {
	val  T
	prev *dListNode[T]
	next *dListNode[T]
}

func NewDListNode[T any](val T, prev *dListNode[T], next *dListNode[T]) *dListNode[T] {
	return &dListNode[T]{val: val, prev: prev, next: next}
}

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}
