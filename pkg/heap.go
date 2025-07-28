package pkg

type Heap[T any] struct {
	data    []T
	compare func(x T, y T) bool
}

func NewHeap[T any](data []T, compare func(x T, y T) bool) *Heap[T] {
	return &Heap[T]{data, compare}
}

func (h *Heap[T]) Len() int           { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool { return h.compare(h.data[i], h.data[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap[T]) Push(x any)         { h.data = append(h.data, x.(T)) }

func (h *Heap[T]) Pop() any { last := h.data[h.Len()-1]; h.data = h.data[:h.Len()-1]; return last }
func (h *Heap[T]) Peak() T {
	if h.Len() > 0 {
		return h.data[0]
	}

	return *new(T)
}

// type hp struct{ sort.IntSlice }
// func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
// func (h *hp) Push(x any)         { h.IntSlice = append(h.IntSlice, x.(int)) }
// func (h *hp) Pop() any           { v := h.IntSlice[h.IntSlice.Len()-1]; h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1]; return v }
