package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMinimumIndex$
func TestMinimumIndex(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		index int
	}{
		{nums: []int{1, 2, 2, 2}, index: 2},
		{nums: []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}, index: 4},
		{nums: []int{3, 3, 3, 3, 7, 2, 2}, index: -1},
	} {
		index := minimumIndex(tc.nums)
		assert.Equal(t, tc.index, index)
	}
}
