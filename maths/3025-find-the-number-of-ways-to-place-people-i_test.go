package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestNumberOfPairs$
func TestNumberOfPairs(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		count  int
	}{
		{points: [][]int{{1, 1}, {2, 2}, {3, 3}}, count: 0},
		{points: [][]int{{6, 2}, {4, 4}, {2, 6}}, count: 2},
		{points: [][]int{{3, 1}, {1, 3}, {1, 1}}, count: 2},
	} {
		count := numberOfPairs(tc.points)
		assert.Equal(t, tc.count, count)
	}
}
