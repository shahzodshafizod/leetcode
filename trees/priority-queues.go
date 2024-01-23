package trees

import "fmt"

type Heap interface {
	Size() int
	IsEmpty() bool
	Peek() int
	Push(int)
	Pop() int
	Println(string)
}

type priorityQueue []int

func NewPriorityQueue() Heap {
	var pq priorityQueue = make([]int, 0)
	return &pq
}

func (p priorityQueue) Println(label string) {
	fmt.Println(label, p)
}

func (p priorityQueue) Size() int {
	return len(p)
}

func (p priorityQueue) IsEmpty() bool {
	return p.Size() == 0
}

func (p priorityQueue) compare(i, j int) bool {
	return p[i] < p[j]
}

func (p *priorityQueue) swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p priorityQueue) parent(idx int) int {
	return (idx - 1) / 2
}

func (p priorityQueue) leftChild(idx int) int {
	return 2*idx + 1
}

func (p priorityQueue) rightChild(idx int) int {
	return 2*idx + 2
}

func (p priorityQueue) Peek() int {
	var elem int
	if !p.IsEmpty() {
		elem = p[0]
	}
	return elem
}

func (p *priorityQueue) Push(x int) {
	*p = append(*p, x)
	p.siftUp()
}

func (p *priorityQueue) siftUp() {
	index := p.Size() - 1
	parent := p.parent(index)
	for parent >= 0 && p.compare(parent, index) {
		p.swap(index, parent)
		index = parent
		parent = p.parent(parent)
	}
}

func (p *priorityQueue) Pop() int {
	if p.IsEmpty() {
		return 0
	}
	var lastElem = (*p)[0]
	p.swap(0, p.Size()-1)
	*p = (*p)[:p.Size()-1]
	p.siftDown()
	return lastElem
}

func (p *priorityQueue) siftDown() {
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
