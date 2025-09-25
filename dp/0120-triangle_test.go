package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinimumTotal$
func TestMinimumTotal(t *testing.T) {
	for _, tc := range []struct {
		triangle [][]int
		total    int
	}{
		{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, total: 11},
		{triangle: [][]int{{-10}}, total: -10},
	} {
		total := minimumTotal(tc.triangle)
		assert.Equal(t, tc.total, total)
	}
}
