package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinIncrementForUnique$
func TestMinIncrementForUnique(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		moves int
	}{
		{nums: []int{1, 2, 2}, moves: 1},
		{nums: []int{3, 2, 1, 2, 1, 7}, moves: 6},
	} {
		moves := minIncrementForUnique(tc.nums)
		assert.Equal(t, tc.moves, moves)
	}
}
