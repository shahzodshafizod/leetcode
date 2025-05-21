package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumberWays$
func TestNumberWays(t *testing.T) {
	for _, tc := range []struct {
		hats [][]int
		ways int
	}{
		{hats: [][]int{{3, 4}, {4, 5}, {5}}, ways: 1},
		{hats: [][]int{{3, 5, 1}, {3, 5}}, ways: 4},
		{hats: [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}, ways: 24},
	} {
		ways := numberWays(tc.hats)
		assert.Equal(t, tc.ways, ways)
	}
}
