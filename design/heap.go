package design

type Heap[T comparable] struct {
	array   []T
	compare func(x T, y T) bool
}

func NewHeap[T comparable](array []T, compare func(x T, y T) bool) *Heap[T] {
	var newArray = make([]T, len(array))
	copy(newArray, array)
	return &Heap[T]{newArray, compare}
}
func (h *Heap[T]) Len() int            { return len(h.array) }
func (h *Heap[T]) Less(i, j int) bool  { return h.compare(h.array[i], h.array[j]) }
func (h *Heap[T]) Swap(i, j int)       { h.array[i], h.array[j] = h.array[j], h.array[i] }
func (h *Heap[T]) Push(x any)          { h.array = append(h.array, x.(T)) }
func (h *Heap[T]) PushMany(array ...T) { h.array = append(h.array, array...) }
func (h *Heap[T]) Pop() any            { last := h.array[h.Len()-1]; h.array = h.array[:h.Len()-1]; return last }
func (h *Heap[T]) Peek() T             { return h.array[0] }
