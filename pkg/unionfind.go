package pkg

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

type UnionFind interface {
	Union(int, int) bool
	Find(int) int
	Reset(int)
}

type unionFind struct {
	parent []int
}

func (u *unionFind) Reset(x int) {
	u.parent[x] = x
}

func NewUnionFind(size int) UnionFind {
	uf := &unionFind{parent: make([]int, size)}
	for i := 0; i < size; i++ {
		// each node's root is the node itself,
		// so nodes are disconnected
		uf.Reset(i)
	}
	return uf
}

func (u *unionFind) Union(x int, y int) bool {
	px := u.Find(x)
	py := u.Find(y)
	if px == py {
		return false
	}
	u.parent[py] = px
	return true
}

func (u *unionFind) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

// QUICK FIND /////////////////////////////////////////////////////////////////

type unionFindQuickFind struct {
	root []int
}

func NewUnionFindQuickFind(size int) UnionFind {
	uf := &unionFindQuickFind{make([]int, size)}
	for idx := 0; idx < size; idx++ {
		uf.Reset(idx)
	}
	return uf
}

// O(N)
func (u *unionFindQuickFind) Union(x int, y int) bool {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return false
	}
	for i := range u.root {
		if u.root[i] == rootY {
			u.root[i] = rootX
		}
	}
	return true
}

// O(1)
func (u *unionFindQuickFind) Find(x int) int {
	return u.root[x]
}

func (u *unionFindQuickFind) Reset(idx int) {
	u.root[idx] = idx
}

// QUICK UNION ////////////////////////////////////////////////////////////////

type unionFindQuickUnion struct {
	root []int
}

func NewUnionFindQuickUnion(size int) UnionFind {
	uf := &unionFindQuickUnion{make([]int, size)}
	for idx := 0; idx < size; idx++ {
		uf.Reset(idx)
	}
	return uf
}

// O(1)
func (u *unionFindQuickUnion) Union(x int, y int) bool {
	rootX := u.root[x]
	rootY := u.root[y]
	if rootX == rootY {
		return false
	}
	u.root[rootY] = rootX
	return true
}

// O(N)
func (u *unionFindQuickUnion) Find(x int) int {
	if x != u.root[x] {
		u.root[x] = u.Find(u.root[x])
	}
	return u.root[x]
}

func (u *unionFindQuickUnion) Reset(idx int) {
	u.root[idx] = idx
}

// RANK UNION /////////////////////////////////////////////////////////////////

type unionFindRanked struct {
	root []int
	rank []int // tree hight
}

func NewUnionFindRanked(size int) UnionFind {
	uf := &unionFindRanked{
		root: make([]int, size),
		rank: make([]int, size),
	}
	for idx := 0; idx < size; idx++ {
		uf.Reset(idx)
	}
	return uf
}

// O(Log N)
func (u *unionFindRanked) Find(x int) int {
	if x != u.root[x] {
		u.root[x] = u.Find(u.root[x])
	}
	return u.root[x]
}

// O(Log N)
func (u *unionFindRanked) Union(x int, y int) bool {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return false
	}
	if u.rank[rootX] < u.rank[rootY] {
		u.root[rootX] = rootY
		u.rank[rootY] += u.rank[rootX]
	} else {
		u.root[rootY] = rootX
		u.rank[rootX] += u.rank[rootY]
	}
	return true
}

func (u *unionFindRanked) Reset(idx int) {
	u.root[idx] = idx
	u.rank[idx] = 1
}
