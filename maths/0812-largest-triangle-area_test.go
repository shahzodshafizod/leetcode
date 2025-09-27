package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestLargestTriangleArea$
func TestLargestTriangleArea(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		area   float64
	}{
		{points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}, area: 2.00000},
		{points: [][]int{{1, 0}, {0, 0}, {0, 1}}, area: 0.50000},
	} {
		area := largestTriangleArea(tc.points)
		assert.InEpsilon(t, tc.area, area, 0.00001)
	}
}
