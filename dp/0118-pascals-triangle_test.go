package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestGenerate$
func TestGenerate(t *testing.T) {
	for _, tc := range []struct {
		numRows int
		rows    [][]int
	}{
		{numRows: 1, rows: [][]int{{1}}},
		{numRows: 5, rows: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{numRows: 6, rows: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}}},
	} {
		rows := generate(tc.numRows)
		assert.Equal(t, tc.rows, rows)
	}
}
