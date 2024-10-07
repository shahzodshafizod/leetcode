package unionfinds

import (
	"sort"

	"github.com/shahzodshafizod/leetcode/design"
)

// https://leetcode.com/problems/rank-transform-of-a-matrix/

/*
Disjoint Set (Union-Find Algorithm)
Two sets are called disjoint sets if they don’t have any element in common,
the intersection of sets is a null set.
*/

func matrixRankTransform(matrix [][]int) [][]int {
	var m = len(matrix)
	var n = len(matrix[0])
	var list = make(map[int][][2]int)
	var answer = make([][]int, m)
	for row := 0; row < m; row++ {
		answer[row] = make([]int, n)
		for col := 0; col < n; col++ {
			list[matrix[row][col]] = append(list[matrix[row][col]], [2]int{row, col})
		}
	}
	var vals = make([]int, 0, len(list))
	for val := range list {
		vals = append(vals, val)
	}
	sort.Ints(vals)
	var rowRanks = make([]int, m)
	var colRanks = make([]int, n)
	var groups = design.NewDisjointSet(m + n)
	var ranks = make(map[int]int)
	for _, val := range vals {
		var cells = list[val]
		for _, cell := range cells {
			groups.Union(cell[0], cell[1]+m)
		}
		for _, cell := range cells {
			groupID := groups.Find(cell[0])
			rank := max(rowRanks[cell[0]], colRanks[cell[1]])
			if rank > ranks[groupID] {
				ranks[groupID] = rank
			}
		}
		for _, cell := range cells {
			rank := ranks[groups.Find(cell[0])] + 1
			rowRanks[cell[0]] = rank
			colRanks[cell[1]] = rank
			answer[cell[0]][cell[1]] = rank
		}
		for _, cell := range cells {
			groups.Reset(cell[0])
			groups.Reset(cell[1] + m)
		}
	}
	return answer
}
