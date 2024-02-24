package design

// UnionFind
type UF interface {
	Union(int, int)
	Find(int) int
	Connected(int, int) bool
	Reset(int)
}

type uf struct {
	parent []int
}

func (u *uf) Reset(x int) {
	u.parent[x] = x
}

func NewUF(size int) UF {
	var u = &uf{parent: make([]int, size)}
	for i := 0; i < size; i++ {
		u.Reset(i)
	}
	return u
}

func (u *uf) Union(x int, y int) {
	u.parent[u.Find(y)] = u.Find(x)
}

func (u *uf) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *uf) Connected(x int, y int) bool {
	return u.Find(x) == u.Find(y)
}
