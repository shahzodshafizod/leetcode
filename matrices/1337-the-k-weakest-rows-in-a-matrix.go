package matrices

import "sort"

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	soldiers := make([]int, m)
	indices := make([]int, m)

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			soldiers[row] += mat[row][col]
		}

		indices[row] = row
	}

	sort.Slice(indices, func(i int, j int) bool {
		if soldiers[indices[i]] == soldiers[indices[j]] {
			return indices[i] < indices[j]
		}

		return soldiers[indices[i]] < soldiers[indices[j]]
	})

	return indices[:k]
}
