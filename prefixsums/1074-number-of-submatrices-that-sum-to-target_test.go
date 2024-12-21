package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestNumSubmatrixSumTarget$
func TestNumSubmatrixSumTarget(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		target int
		count  int
	}{
		{matrix: [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, target: 0, count: 4},
		{matrix: [][]int{{1, -1}, {-1, 1}}, target: 0, count: 5},
		{matrix: [][]int{{904}}, target: 0, count: 0},
	} {
		count := numSubmatrixSumTarget(tc.matrix, tc.target)
		assert.Equal(t, tc.count, count)
	}
}
