package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestAreaOfMaxDiagonal$
func TestAreaOfMaxDiagonal(t *testing.T) {
	for _, tc := range []struct {
		dimensions [][]int
		area       int
	}{
		{dimensions: [][]int{{9, 3}, {8, 6}}, area: 48},
		{dimensions: [][]int{{3, 4}, {4, 3}}, area: 12},
	} {
		area := areaOfMaxDiagonal(tc.dimensions)
		assert.Equal(t, tc.area, area)
	}
}
