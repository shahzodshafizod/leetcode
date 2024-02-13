package design

type UnionFind interface {
	Union(int, int)
	Find(int) int
	Reset(int)
}

type unionfind []int

func (uf *unionfind) Reset(x int) {
	(*uf)[x] = x
}

func NewUnionFind(size int) UnionFind {
	var uf = unionfind(make([]int, size))
	for i := 0; i < size; i++ {
		uf.Reset(i)
	}
	return &uf
}

func (uf *unionfind) Union(x int, y int) {
	(*uf)[uf.Find(y)] = uf.Find(x)
}

func (uf *unionfind) Find(x int) int {
	if (*uf)[x] != x {
		(*uf)[x] = uf.Find((*uf)[x])
	}
	return (*uf)[x]
}
