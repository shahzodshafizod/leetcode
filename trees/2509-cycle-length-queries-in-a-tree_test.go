package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestCycleLengthQueries$
func TestCycleLengthQueries(t *testing.T) {
	for _, tc := range []struct {
		n       int
		queries [][]int
		result  []int
	}{
		{
			n:       3,
			queries: [][]int{{5, 3}, {4, 7}, {2, 3}},
			result:  []int{4, 5, 3},
		},
		{
			n:       2,
			queries: [][]int{{1, 2}},
			result:  []int{2},
		},
		{
			n:       5,
			queries: [][]int{{8, 5}, {14, 10}},
			result:  []int{4, 7},
		},
		{
			n:       4,
			queries: [][]int{{7, 2}, {15, 1}},
			result:  []int{4, 4},
		},
	} {
		result := cycleLengthQueries(tc.n, tc.queries)
		assert.Equal(t, tc.result, result)
	}
}
