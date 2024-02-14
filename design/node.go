package design

type node[T any] struct {
	val  T
	next *node[T]
}

type listNode[T any] struct {
	val  T
	next *listNode[T]
	prev *listNode[T]
}
