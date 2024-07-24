package design

type PQSort interface {
	Sort()
}

type pqSort struct {
	length  int
	compare func(i int, j int) bool
	swap    func(i int, j int)
}

func NewPQSort(length int, compare func(i int, j int) bool, swap func(i int, j int)) PQSort {
	var pq = &pqSort{
		length:  length,
		compare: compare,
		swap:    swap,
	}
	pq.heapify()
	return pq
}

func (p *pqSort) heapify() {
	for parent := p.length/2 - 1; parent >= 0; parent-- {
		p.shiftDown(parent, p.length)
	}
}

func (p *pqSort) Sort() {
	var n = p.length
	for n > 0 {
		n--
		p.swap(0, n)
		p.shiftDown(0, n)
	}
}

func (p *pqSort) left(parent int) int {
	return 2*parent + 1 // right: 2*parent + 2
}

func (p *pqSort) shiftDown(parent int, n int) {
	var child = p.left(parent)
	for child < n {
		if child+1 < n && p.compare(child, child+1) {
			child++
		}
		if !p.compare(parent, child) {
			break
		}
		p.swap(parent, child)
		parent = child
		child = p.left(parent)
	}
}
