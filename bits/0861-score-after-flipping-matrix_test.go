package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMatrixScore$
func TestMatrixScore(t *testing.T) {
	for _, tc := range []struct {
		grid  [][]int
		score int
	}{
		{grid: [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}, score: 39},
		{grid: [][]int{{0}}, score: 1},
	} {
		score := matrixScore(tc.grid)
		assert.Equal(t, tc.score, score)
	}
}
