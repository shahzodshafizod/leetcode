package design

type PriorityQueue[T comparable] interface {
	Size() int
	IsEmpty() bool
	Peek() T
	Push(T)
	Pop() T
}

type priorityQueue[T comparable] struct {
	array      []T // binary tree as array
	comparator func(item1 T, item2 T) bool
}

func NewPriorityQueue[T comparable](comparator func(item1 T, item2 T) bool) PriorityQueue[T] {
	var pq = priorityQueue[T]{
		array:      make([]T, 0),
		comparator: comparator,
	}
	return &pq
}

func (p priorityQueue[T]) Size() int {
	return len(p.array)
}

func (p priorityQueue[T]) IsEmpty() bool {
	return p.Size() == 0
}

func (p priorityQueue[T]) compare(i int, j int) bool {
	return p.comparator(p.array[i], p.array[j])
}

func (p *priorityQueue[T]) swap(i int, j int) {
	(*p).array[i], (*p).array[j] = (*p).array[j], (*p).array[i]
}

func (p priorityQueue[T]) parent(idx int) int {
	return (idx - 1) / 2
}

func (p priorityQueue[T]) leftChild(idx int) int {
	return 2*idx + 1
}

func (p priorityQueue[T]) rightChild(idx int) int {
	return 2*idx + 2
}

func (p priorityQueue[T]) Peek() T {
	var item T
	if !p.IsEmpty() {
		item = p.array[0]
	}
	return item
}

func (p *priorityQueue[T]) Push(x T) {
	(*p).array = append((*p).array, x)
	p.siftUp()
}

func (p *priorityQueue[T]) siftUp() {
	index := p.Size() - 1
	parent := p.parent(index)
	for parent >= 0 && p.compare(parent, index) {
		p.swap(index, parent)
		index = parent
		parent = p.parent(parent)
	}
}

func (p *priorityQueue[T]) Pop() T {
	var lastItem T
	if p.IsEmpty() {
		return lastItem
	}
	lastItem = (*p).array[0]
	p.swap(0, p.Size()-1)
	(*p).array = (*p).array[:p.Size()-1]
	p.siftDown()
	return lastItem
}

func (p *priorityQueue[T]) siftDown() {
	parent := 0

	for p.leftChild(parent) < p.Size() && p.compare(parent, p.leftChild(parent)) ||
		p.rightChild(parent) < p.Size() && p.compare(parent, p.rightChild(parent)) {

		if p.rightChild(parent) < p.Size() && p.compare(parent, p.rightChild(parent)) {

			if p.compare(parent, p.leftChild(parent)) { // both

				if p.compare(p.leftChild(parent), p.rightChild(parent)) {
					p.swap(parent, p.rightChild(parent))
					parent = p.rightChild(parent)
				} else {
					p.swap(parent, p.leftChild(parent))
					parent = p.leftChild(parent)
				}

			} else { // just right
				p.swap(parent, p.rightChild(parent))
				parent = p.rightChild(parent)
			}

		} else { // just left
			p.swap(parent, p.leftChild(parent))
			parent = p.leftChild(parent)
		}
	}
}
