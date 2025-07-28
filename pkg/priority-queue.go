package pkg

type priorityQueue[T any] interface {
	Heapify()
	Len() int
	Peek() T
	Push(T)
	Pop() T
}

type pq[T any] struct {
	data []T // binary tree as array
	less func(x T, y T) bool
}

func newPriorityQueue[T any](data []T, less func(x T, y T) bool) priorityQueue[T] {
	pq := pq[T]{
		data: data,
		less: less,
	}

	return &pq
}

func (p *pq[T]) Heapify() { // O(NLogN)
	for parent := p.Len()/2 - 1; parent >= 0; parent-- {
		p.shiftDown(parent, p.Len())
	}
}

func (p pq[T]) Len() int {
	return len(p.data)
}

func (p pq[T]) Peek() T {
	var val T
	if p.Len() > 0 {
		val = p.data[0]
	}

	return val
}

func (p *pq[T]) Push(x T) {
	p.data = append(p.data, x)
	p.shiftUp(p.Len() - 1)
}

func (p *pq[T]) Pop() T {
	var root T
	if p.Len() == 0 {
		return root
	}

	root = p.data[0]
	p.swap(0, p.Len()-1)
	p.data = p.data[:p.Len()-1]
	p.shiftDown(0, p.Len())

	return root
}

func (p pq[T]) compare(i int, j int) bool { return p.less(p.data[i], p.data[j]) }
func (p *pq[T]) swap(i int, j int)        { p.data[i], p.data[j] = p.data[j], p.data[i] }
func (p pq[T]) parent(idx int) int        { return (idx - 1) / 2 }
func (p pq[T]) left(idx int) int          { return 2*idx + 1 } // right: 2*idx + 2

func (p *pq[T]) shiftUp(child int) {
	parent := p.parent(child)
	for parent >= 0 && p.compare(parent, child) {
		p.swap(parent, child)
		child = parent
		parent = p.parent(child)
	}
}

func (p *pq[T]) shiftDown(lo int, hi int) {
	parent := lo

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
