package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMaxPoints$
func TestMaxPoints(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		count  int
	}{
		{points: [][]int{{1, 1}, {2, 2}, {3, 3}}, count: 3},
		{points: [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, count: 4},
		{points: [][]int{{0, 0}}, count: 1},
		{points: [][]int{{0, 0}, {94911151, 94911150}, {94911152, 94911151}}, count: 3},
		{points: [][]int{{0, 0}, {1, 1}, {0, 0}}, count: 2},
		{points: [][]int{{0, 0}, {0, 0}}, count: 2},
		{points: [][]int{{0, 0}, {1073741822, 2147483645}, {1073741823, 2147483647}}, count: 3},
		{points: [][]int{{1, 1}, {1, 1}, {2, 3}}, count: 2},
	} {
		count := maxPoints(tc.points)
		assert.Equal(t, tc.count, count)
	}
}
