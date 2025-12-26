package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCountOfPairs$
func TestCountOfPairs(t *testing.T) {
	for _, tc := range []struct {
		n      int
		x      int
		y      int
		result []int64
	}{
		{n: 3, x: 1, y: 3, result: []int64{6, 0, 0}},
		{n: 5, x: 2, y: 4, result: []int64{10, 8, 2, 0, 0}},
		{n: 4, x: 1, y: 1, result: []int64{6, 4, 2, 0}},
	} {
		result := countOfPairs(tc.n, tc.x, tc.y)
		assert.Equal(t, tc.result, result)
	}
}
