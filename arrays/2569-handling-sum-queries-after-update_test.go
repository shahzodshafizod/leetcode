package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestHandleQuery$
func TestHandleQuery(t *testing.T) {
	for _, tc := range []struct {
		nums1   []int
		nums2   []int
		queries [][]int
		output  []int64
	}{
		{
			nums1:   []int{1, 0, 1},
			nums2:   []int{0, 0, 0},
			queries: [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}},
			output:  []int64{3},
		},
		{
			nums1:   []int{1},
			nums2:   []int{5},
			queries: [][]int{{2, 0, 0}, {3, 0, 0}},
			output:  []int64{5},
		},
		{
			nums1:   []int{1, 0, 1, 0, 1},
			nums2:   []int{1, 2, 3, 4, 5},
			queries: [][]int{{1, 0, 2}, {2, 1, 0}, {3, 0, 0}, {1, 1, 3}, {2, 2, 0}, {3, 0, 0}},
			output:  []int64{17, 23},
		},
		{
			nums1:   []int{0, 0, 0},
			nums2:   []int{1, 2, 3},
			queries: [][]int{{1, 0, 2}, {2, 3, 0}, {3, 0, 0}, {1, 1, 1}, {2, 2, 0}, {3, 0, 0}},
			output:  []int64{15, 19},
		},
	} {
		output := handleQuery(tc.nums1, tc.nums2, tc.queries)
		assert.Equal(t, tc.output, output)
	}
}
