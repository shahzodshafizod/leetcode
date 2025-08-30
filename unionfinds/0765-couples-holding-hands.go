package unionfinds

/*
Permutation graph

The answer to the minimum number of swaps is
the initial number of couples n minus the number of connected components
after the union-find process.
*/

// https://leetcode.com/problems/couples-holding-hands/

func minSwapsCouples(row []int) int {
	// bitwise right shift operator to divide by 2
	n := len(row) >> 1 // we have len(row)/2 couples
	p := make([]int, n)

	for i := range n {
		p[i] = i
	}

	connected := n

	var find func(x int) int

	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}

		return p[x]
	}
	union := func(x int, y int) {
		px, py := find(x), find(y)
		if px != py {
			p[px] = py
			connected--
		}
	}

	for i := range n {
		union(row[2*i]>>1, row[2*i+1]>>1)
		// OR union(row[2*i]/2, row[2*i+1]/2)
	}

	return n - connected
}
