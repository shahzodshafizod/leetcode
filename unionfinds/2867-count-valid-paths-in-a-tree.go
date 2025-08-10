package unionfinds

import (
	"math"
)

// https://leetcode.com/problems/count-valid-paths-in-a-tree/

// Approach: DSU (Disjoint Set Union)
// Time: O(n log n)
// Space: O(n)
func countPaths2867(n int, edges [][]int) int64 {
	primes := make([]bool, n+1)

	roots := make([]int, n+1)
	for node := 1; node <= n; node++ {
		primes[node] = true
		roots[node] = node
	}

	primes[0], primes[1] = false, false
	for num, sqrt := 2, int(math.Sqrt(float64(n))); num <= sqrt; num++ {
		if primes[num] {
			for nxt := 2 * num; nxt <= n; nxt += num {
				primes[nxt] = false
			}
		}
	}

	var find func(int) int

	find = func(x int) int {
		if roots[x] != x {
			roots[x] = find(roots[x])
		}

		return roots[x]
	}

	union := func(u int, v int) {
		ru := find(u)
		rv := find(v)
		roots[min(ru, rv)] = max(ru, rv)
	}

	graph := make([][]int, n+1)

	var u, v int
	for _, edge := range edges {
		u, v = edge[0], edge[1]
		graph[u] = append(graph[u], v)

		graph[v] = append(graph[v], u)
		if !primes[u] && !primes[v] {
			union(u, v)
		}
	}

	sizes := make([]int64, n+1)
	for node := 1; node <= n; node++ {
		if !primes[node] {
			sizes[find(node)]++
		}
	}

	var res, total int64

	for node := 1; node <= n; node++ {
		if primes[node] {
			total = 0

			for _, nxt := range graph[node] {
				if !primes[nxt] {
					total += sizes[find(nxt)]
				}
			}

			res += total

			for _, nxt := range graph[node] {
				if !primes[nxt] {
					total -= sizes[find(nxt)]
					res += sizes[find(nxt)] * total
				}
			}
		}
	}

	return res
}
