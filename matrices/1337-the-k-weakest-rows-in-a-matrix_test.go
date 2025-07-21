package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestKWeakestRows$
func TestKWeakestRows(t *testing.T) {
	for _, tc := range []struct {
		mat     [][]int
		k       int
		indices []int
	}{
		{mat: [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, k: 3, indices: []int{2, 0, 3}},
		{mat: [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, k: 2, indices: []int{0, 2}},
	} {
		indices := kWeakestRows(tc.mat, tc.k)
		assert.Equal(t, tc.indices, indices)
	}
}
