package design

/*
The main idea of a “disjoint set” is to have all connected vertices have
the same parent node or root node, whether directly or indirectly connected.
To check if two vertices are connected,
we only need to check if they have the same root node.

There are two ways to implement a “disjoint set”.
	- Implementation with Quick Find:
		in this case, the time complexity of the find function will be O(1).
		However, the union function will take more time with the time
		complexity of O(N).
	- Implementation with Quick Union:
		compared with the Quick Find implementation, the time complexity of the
		union function is better. Meanwhile, the find function will take more
		time in this case.
*/

type DisjointSet interface {
	Union(int, int)
	Find(int) int
	Connected(int, int) bool
	Reset(int)
}

type disjointSet struct {
	parent []int
}

func (d *disjointSet) Reset(x int) {
	d.parent[x] = x
}

func NewDisjointSet(size int) DisjointSet {
	var d = &disjointSet{parent: make([]int, size)}
	for i := 0; i < size; i++ {
		// each node's root is the node itself,
		// so nodes are disconnected
		d.Reset(i)
	}
	return d
}

func (d *disjointSet) Union(x int, y int) {
	d.parent[d.Find(y)] = d.Find(x)
}

func (d *disjointSet) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *disjointSet) Connected(x int, y int) bool {
	return d.Find(x) == d.Find(y)
}

// QUICK FIND /////////////////////////////////////////////////////////////////

type quickFind struct {
	root []int
}

func NewQuickFind(size int) DisjointSet {
	var quickFind = &quickFind{make([]int, size)}
	for idx := range quickFind.root {
		quickFind.Reset(idx)
	}
	return quickFind
}

// O(N)
func (f *quickFind) Union(x int, y int) {
	var rootX = f.Find(x)
	var rootY = f.Find(y)
	if rootX != rootY {
		for i := range f.root {
			if f.root[i] == rootY {
				f.root[i] = rootX
			}
		}
	}
}

// O(1)
func (f *quickFind) Find(x int) int {
	return f.root[x]
}

// O(1)
func (f *quickFind) Connected(x int, y int) bool {
	return f.Find(x) == f.Find(y)
}

func (f *quickFind) Reset(idx int) {
	f.root[idx] = idx
}

// QUICK UNION ////////////////////////////////////////////////////////////////

type quickUnion struct {
	root []int
}

func NewQuickUnion(size int) DisjointSet {
	var quickUnion = &quickUnion{make([]int, size)}
	for idx := range quickUnion.root {
		quickUnion.Reset(idx)
	}
	return quickUnion
}

// O(N)
func (u *quickUnion) Union(x int, y int) {
	var rootX = u.Find(x) // O(N)
	var rootY = u.Find(y) // O(N)
	if rootX != rootY {
		u.root[rootY] = rootX
	}
}

// O(N)
func (u *quickUnion) Find(x int) int {
	for x != u.root[x] {
		x = u.root[x] // optimize?
	}
	return x
}

// O(N)
func (u *quickUnion) Connected(x int, y int) bool {
	return u.Find(x) == u.Find(y)
}

func (u *quickUnion) Reset(idx int) {
	u.root[idx] = idx
}

// RANK UNION /////////////////////////////////////////////////////////////////

type rankUnion struct {
	root []int
	rank []int // tree hight
}

func NewRankUnion(size int) DisjointSet {
	var rankUnion = &rankUnion{
		root: make([]int, size),
		rank: make([]int, size),
	}
	for idx := range rankUnion.root {
		rankUnion.Reset(idx)
	}
	return rankUnion
}

// O(Log N)
func (r *rankUnion) Union(x int, y int) {
	var rootX = r.Find(x)
	var rootY = r.Find(y)
	if rootX != rootY {
		if r.rank[rootX] > r.rank[rootY] {
			r.root[rootY] = rootX
		} else if r.rank[rootX] < r.rank[rootY] {
			r.root[rootX] = rootY
		} else {
			r.root[rootY] = rootX
			r.rank[rootX]++
		}
	}
}

// O(Log N)
func (r *rankUnion) Find(x int) int {
	for x != r.root[x] {
		x = r.root[x]
	}
	return x
}

// O(Log N)
func (r *rankUnion) Connected(x int, y int) bool {
	return r.Find(x) == r.Find(y)
}

func (r *rankUnion) Reset(idx int) {
	r.root[idx] = idx
	r.rank[idx] = 1
}

// OPTIMIZED DISJOINT SET /////////////////////////////////////////////////////

type optimizedDS struct {
	root []int
	rank []int
}

func NewOptimizedDS(size int) DisjointSet {
	var optimized = &optimizedDS{
		root: make([]int, size),
		rank: make([]int, size),
	}
	for idx := range optimized.root {
		optimized.Reset(idx)
	}
	return optimized
}

func (o *optimizedDS) Find(x int) int {
	if x != o.root[x] {
		o.root[x] = o.Find(o.root[x])
	}
	return o.root[x]
}

func (o *optimizedDS) Union(x int, y int) {
	rootX := o.Find(x)
	rootY := o.Find(y)
	if rootX != rootY {
		if o.rank[rootX] > o.rank[rootY] {
			o.root[rootY] = rootX
		} else if o.rank[rootX] < o.rank[rootY] {
			o.root[rootX] = rootY
		} else {
			o.root[rootY] = rootX
			o.rank[rootX]++
		}
	}
}

func (o *optimizedDS) Connected(x int, y int) bool {
	return o.Find(x) == o.Find(y)
}

func (o *optimizedDS) Reset(idx int) {
	o.root[idx] = idx
	o.rank[idx] = 1
}
