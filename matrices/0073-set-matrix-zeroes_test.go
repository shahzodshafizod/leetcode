package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestSetZeroes$
func TestSetZeroes(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		result [][]int
	}{
		{matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, result: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, result: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	} {
		setZeroes(tc.matrix)
		assert.Equal(t, tc.result, tc.matrix)
	}
}
