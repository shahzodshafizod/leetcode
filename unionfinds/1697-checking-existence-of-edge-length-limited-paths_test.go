package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestDistanceLimitedPathsExist$
func TestDistanceLimitedPathsExist(t *testing.T) {
	for _, tc := range []struct {
		n        int
		edgeList [][]int
		queries  [][]int
		output   []bool
	}{
		{
			n:        3,
			edgeList: [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}},
			queries:  [][]int{{0, 1, 2}, {0, 2, 5}},
			output:   []bool{false, true},
		},
		{
			n:        5,
			edgeList: [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}},
			queries:  [][]int{{0, 4, 14}, {1, 4, 13}},
			output:   []bool{true, false},
		},
	} {
		output := distanceLimitedPathsExist(tc.n, tc.edgeList, tc.queries)
		assert.Equal(t, tc.output, output)
	}
}
