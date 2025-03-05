package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestColoredCells$
func TestColoredCells(t *testing.T) {
	for _, tc := range []struct {
		n     int
		cells int64
	}{
		{n: 1, cells: 1},
		{n: 2, cells: 5},
		{n: 100000, cells: 19999800001},
	} {
		cells := coloredCells(tc.n)
		assert.Equal(t, tc.cells, cells)
	}
}
