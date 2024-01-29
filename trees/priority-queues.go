package trees

type Heap interface {
	Size() int
	IsEmpty() bool
	Peek() int
	Has(int) bool
	Push(int)
	Pop() int
}

type priorityQueue struct {
	array      []int // binary tree as array
	comparator func(i int, j int) bool
}

func NewPriorityQueue(comparator func(i int, j int) bool) Heap {
	var pq = priorityQueue{
		array:      make([]int, 0),
		comparator: comparator,
	}
	return &pq
}

func (p priorityQueue) Size() int {
	return len(p.array)
}

func (p priorityQueue) IsEmpty() bool {
	return p.Size() == 0
}

func (p priorityQueue) compare(i, j int) bool {
	return p.comparator(p.array[i], p.array[j])
}

func (p *priorityQueue) swap(i, j int) {
	(*p).array[i], (*p).array[j] = (*p).array[j], (*p).array[i]
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
	var item int
	if !p.IsEmpty() {
		item = p.array[0]
	}
	return item
}

func (p priorityQueue) Has(v int) bool {
	for _, item := range p.array {
		if item == v {
			return true
		}
	}
	return false
}

func (p *priorityQueue) Push(x int) {
	(*p).array = append((*p).array, x)
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
	var lastItem = (*p).array[0]
	p.swap(0, p.Size()-1)
	(*p).array = (*p).array[:p.Size()-1]
	p.siftDown()
	return lastItem
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
