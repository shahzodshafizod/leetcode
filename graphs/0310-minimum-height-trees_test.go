package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindMinHeightTrees$
func TestFindMinHeightTrees(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		trees map[int]bool
	}{
		{n: 4, edges: [][]int{{1, 0}, {1, 2}, {1, 3}}, trees: map[int]bool{1: true}},
		{n: 6, edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}, trees: map[int]bool{3: true, 4: true}},
		{n: 2, edges: [][]int{{0, 1}}, trees: map[int]bool{1: true, 0: true}},
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}}, trees: map[int]bool{3: true}},
		{n: 1, edges: [][]int{}, trees: map[int]bool{0: true}},
	} {
		roots := findMinHeightTrees(tc.n, tc.edges)
		trees := make(map[int]bool) // because the order is randomly organized

		for _, root := range roots {
			trees[root] = true
		}

		assert.Equal(t, tc.trees, trees)
	}
}
