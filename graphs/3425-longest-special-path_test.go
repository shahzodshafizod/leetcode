package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestLongestSpecialPath$
func TestLongestSpecialPath(t *testing.T) {
	for _, tc := range []struct {
		edges  [][]int
		nums   []int
		result []int
	}{
		{
			edges:  [][]int{{0, 1, 2}, {1, 2, 3}, {1, 3, 5}, {1, 4, 4}, {2, 5, 6}},
			nums:   []int{2, 1, 2, 1, 3, 1},
			result: []int{6, 2},
		},
		{
			edges:  [][]int{{1, 0, 8}},
			nums:   []int{2, 2},
			result: []int{0, 1},
		},
		{
			edges:  [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}},
			nums:   []int{1, 2, 3, 4},
			result: []int{3, 4},
		},
	} {
		result := longestSpecialPath(tc.edges, tc.nums)
		assert.Equal(t, tc.result, result)
	}
}
