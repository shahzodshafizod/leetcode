package unionfinds

// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/

type uf struct {
	n    int
	par  []int
	rank []int
}

func newUF(n int) *uf {
	var uf = &uf{
		n:    n,
		par:  make([]int, n+1),
		rank: make([]int, n+1),
	}
	for i := 0; i < n; i++ {
		uf.par[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (u *uf) find(x int) int {
	for x != u.par[x] {
		u.par[x] = u.par[u.par[x]]
		x = u.par[x]
	}
	return x
}

// Return 1 if success, 0 otherwise
func (u *uf) union(x1 int, x2 int) int {
	p1, p2 := u.find(x1), u.find(x2)
	if p1 == p2 {
		return 0
	}
	if u.rank[p1] > u.rank[p2] {
		u.rank[p1] += u.rank[p2]
		u.par[p2] = p1
	} else {
		u.rank[p2] += u.rank[p1]
		u.par[p1] = p2
	}
	u.n--
	return 1
}

func (u *uf) isConnected() bool {
	return u.n == 1
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice, bob := newUF(n), newUF(n)
	cnt := 0 // num of edges we have keep
	for _, edge := range edges {
		t, src, dst := edge[0], edge[1], edge[2]
		if t == 3 {
			cnt += alice.union(src, dst) | bob.union(src, dst)
		}
	}
	for _, edge := range edges {
		t, src, dst := edge[0], edge[1], edge[2]
		if t == 1 {
			cnt += alice.union(src, dst)
		} else {
			cnt += bob.union(src, dst)
		}
	}
	if alice.isConnected() && bob.isConnected() {
		return len(edges) - cnt
	}
	return -1
}
