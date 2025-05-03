package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinDominoRotations$
func TestMinDominoRotations(t *testing.T) {
	for _, tc := range []struct {
		tops      []int
		bottoms   []int
		rotations int
	}{
		{tops: []int{2, 1, 2, 4, 2, 2}, bottoms: []int{5, 2, 6, 2, 3, 2}, rotations: 2},
		{tops: []int{3, 5, 1, 2, 3}, bottoms: []int{3, 6, 3, 3, 4}, rotations: -1},
		{tops: []int{1, 1, 1, 2, 2, 2}, bottoms: []int{5, 5, 5, 1, 1, 1}, rotations: 3},
	} {
		rotations := minDominoRotations(tc.tops, tc.bottoms)
		assert.Equal(t, tc.rotations, rotations)
	}
}
