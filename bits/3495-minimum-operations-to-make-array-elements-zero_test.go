package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMinOperations3495$
func TestMinOperations3495(t *testing.T) {
	for _, tc := range []struct {
		queries [][]int
		count   int64
	}{
		{queries: [][]int{{1, 2}, {2, 4}}, count: 3},
		{queries: [][]int{{2, 6}}, count: 4},
		{queries: [][]int{{1, 8}}, count: 7},
	} {
		count := minOperations3495(tc.queries)
		assert.Equal(t, tc.count, count)
	}
}
