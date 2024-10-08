package pkg

type DListNode[T any] struct {
	Val  T
	Prev *DListNode[T]
	Next *DListNode[T]
}

func NewDListNode[T any](val T, prev *DListNode[T], next *DListNode[T]) *DListNode[T] {
	return &DListNode[T]{Val: val, Prev: prev, Next: next}
}
