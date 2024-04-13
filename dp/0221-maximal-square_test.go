package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaximalSquare$
func TestMaximalSquare(t *testing.T) {
	for _, tc := range []struct {
		matrix  [][]byte
		maximal int
	}{
		{matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, maximal: 4},
		{matrix: [][]byte{{'0', '1'}, {'1', '0'}}, maximal: 1},
		{matrix: [][]byte{{'0'}}, maximal: 0},
		{matrix: [][]byte{{'1', '1'}}, maximal: 1},
		{matrix: [][]byte{{'1'}}, maximal: 1},
	} {
		maximal := maximalSquare(tc.matrix)
		assert.Equal(t, tc.maximal, maximal)
	}
}
