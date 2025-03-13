package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestMinZeroArray$
func TestMinZeroArray(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		k       int
	}{
		{nums: []int{3, 6, 4}, queries: [][]int{{2, 2, 4}}, k: -1},
		{nums: []int{0, 2, 0, 4}, queries: [][]int{{3, 3, 3}, {0, 0, 5}}, k: -1},
		{nums: []int{5}, queries: [][]int{{0, 0, 5}, {0, 0, 1}, {0, 0, 3}, {0, 0, 2}}, k: 1},
		{nums: []int{0}, queries: [][]int{{0, 0, 2}, {0, 0, 4}, {0, 0, 4}, {0, 0, 3}, {0, 0, 5}}, k: 0},
		{nums: []int{1, 0, 6}, queries: [][]int{{1, 2, 1}, {0, 0, 4}, {1, 1, 5}, {0, 0, 5}, {1, 2, 4}, {0, 2, 2}, {2, 2, 4}, {1, 2, 2}, {1, 2, 4}, {0, 1, 3}}, k: 6},
		{nums: []int{1, 0, 6}, queries: [][]int{{1, 2, 1}, {0, 0, 4}, {1, 1, 5}, {0, 0, 5}, {1, 2, 4}, {0, 2, 2}, {2, 2, 4}, {1, 2, 2}, {1, 2, 4}, {0, 1, 3}}, k: 6},
	} {
		k := minZeroArray(tc.nums, tc.queries)
		assert.Equal(t, tc.k, k)
	}
}
