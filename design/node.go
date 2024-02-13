package design

type node[T any] struct {
	val  T
	next *node[T]
}
