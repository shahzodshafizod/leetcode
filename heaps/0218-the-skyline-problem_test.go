package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestGetSkyline$
func TestGetSkyline(t *testing.T) {
	for _, tc := range []struct {
		buildings [][]int
		output    [][]int
	}{
		{
			[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{
			[][]int{{0, 2, 3}, {2, 5, 3}},
			[][]int{{0, 3}, {5, 0}},
		},
		{
			[][]int{{0, 3, 3}},
			[][]int{{0, 3}, {3, 0}},
		},
		{
			[][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}},
			[][]int{{1, 3}, {2, 0}},
		},
		{
			[][]int{{0, 2, 3}, {0, 2, 3}},
			[][]int{{0, 3}, {2, 0}},
		},
	} {
		output := getSkyline(tc.buildings)
		assert.Equal(t, tc.output, output)
	}
}
