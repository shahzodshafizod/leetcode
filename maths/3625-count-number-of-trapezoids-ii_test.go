package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountTrapezoids$
func TestCountTrapezoids(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		count  int
	}{
		{points: [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}}, count: 2},
		{points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}, count: 1},
		{points: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}, count: 0}, // All collinear
		{points: [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, count: 1}, // Square
	} {
		count := countTrapezoids(tc.points)
		assert.Equal(t, tc.count, count)
	}
}
