package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestHeightChecker$
func TestHeightChecker(t *testing.T) {
	for _, tc := range []struct {
		heights   []int
		misplaced int
	}{
		{heights: []int{1, 1, 4, 2, 1, 3}, misplaced: 3},
		{heights: []int{5, 1, 2, 3, 4}, misplaced: 5},
		{heights: []int{1, 2, 3, 4, 5}, misplaced: 0},
	} {
		misplaced := heightChecker(tc.heights)
		assert.Equal(t, tc.misplaced, misplaced)
	}
}
