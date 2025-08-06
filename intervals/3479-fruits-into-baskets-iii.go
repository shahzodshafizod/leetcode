package intervals

// https://leetcode.com/problems/fruits-into-baskets-iii/

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	segtree := make([]int, 4*n)

	var build func(node int, left int, right int)

	build = func(node int, left int, right int) {
		if left == right {
			segtree[node] = baskets[left]

			return
		}

		mid := left + (right-left)/2
		build(2*node, left, mid)
		build(2*node+1, mid+1, right)
		segtree[node] = max(segtree[2*node], segtree[2*node+1])
	}

	build(1, 0, n-1)

	var find func(node int, fruit int, left int, right int) bool

	find = func(node int, fruit int, left int, right int) bool {
		if fruit > segtree[node] { // failure base case
			return false
		}

		if left == right { // success base case
			segtree[node] = 0

			return true
		}

		mid := left + (right-left)/2

		res := find(2*node, fruit, left, mid)
		if !res {
			res = find(2*node+1, fruit, mid+1, right)
		}

		segtree[node] = max(segtree[2*node], segtree[2*node+1])

		return res
	}

	res := 0

	for _, fruit := range fruits {
		if !find(1, fruit, 0, n-1) {
			res++
		}
	}

	return res
}
