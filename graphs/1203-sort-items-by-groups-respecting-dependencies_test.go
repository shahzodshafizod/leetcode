package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestSortItems$
func TestSortItems(t *testing.T) {
	for _, tc := range []struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
		result      []int
	}{
		{n: 8, m: 2, group: []int{-1, -1, 1, 0, 0, 1, 0, -1}, beforeItems: [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}, result: []int{6, 3, 4, 5, 2, 0, 7, 1}},
		{n: 8, m: 2, group: []int{-1, -1, 1, 0, 0, 1, 0, -1}, beforeItems: [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}}, result: []int{}},
	} {
		result := sortItems(tc.n, tc.m, tc.group, tc.beforeItems)
		assert.Equal(t, tc.result, result)
	}
}
