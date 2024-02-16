package design

type PriorityQueue[T any] interface {
	Len() int
	Peek() T
	Push(T)
	Pop() T
	Sort()
	Array() []T
}

type priorityQueue[T any] struct {
	array []T // binary tree as array
	less  func(x T, y T) bool
}

func NewPriorityQueue[T any](less func(x T, y T) bool) PriorityQueue[T] {
	var pq = priorityQueue[T]{
		array: make([]T, 0),
		less:  less,
	}
	return &pq
}

func (p priorityQueue[T]) Len() int {
	return len(p.array)
}

func (p priorityQueue[T]) Peek() T {
	var val T
	if p.Len() > 0 {
		val = p.array[0]
	}
	return val
}

func (p *priorityQueue[T]) Push(x T) {
	(*p).array = append((*p).array, x)
	p.siftUp()
}

func (p *priorityQueue[T]) Pop() T {
	var last T
	if p.Len() == 0 {
		return last
	}
	p.swap(0, p.Len()-1)
	last = (*p).array[p.Len()-1]
	(*p).array = (*p).array[:p.Len()-1]
	p.siftDown(p.Len())
	return last
}

func (p *priorityQueue[T]) Sort() {
	var len = p.Len()
	for len > 0 {
		len--
		p.swap(0, len)
		p.siftDown(len)
	}
}

func (p *priorityQueue[T]) Array() []T {
	return p.array
}

func (p priorityQueue[T]) compare(i int, j int) bool {
	return p.less(p.array[i], p.array[j])
}

func (p *priorityQueue[T]) swap(i int, j int) {
	(*p).array[i], (*p).array[j] = (*p).array[j], (*p).array[i]
}

func (p priorityQueue[T]) parent(idx int) int {
	return (idx - 1) / 2
}

func (p priorityQueue[T]) left(idx int) int {
	return 2*idx + 1 // right: 2*idx + 2
}

func (p *priorityQueue[T]) siftUp() {
	child := p.Len() - 1
	var parent = p.parent(child)
	for parent >= 0 && p.compare(parent, child) {
		p.swap(child, parent)
		child = parent
		parent = p.parent(child)
	}
}

func (p *priorityQueue[T]) siftDown(hi int) {
	parent := 0
	child := p.left(parent)
	for child < hi {
		if child+1 < hi && p.compare(child, child+1) { // right
			child++
		}
		if !p.compare(parent, child) {
			return
		}
		p.swap(parent, child)
		parent = child
		child = p.left(parent)
	}
}
