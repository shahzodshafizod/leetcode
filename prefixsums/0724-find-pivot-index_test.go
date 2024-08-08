package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestPivotIndex$
func TestPivotIndex(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		index int
	}{
		{nums: []int{1, 7, 3, 6, 5, 6}, index: 3},
		{nums: []int{1, 2, 3}, index: -1},
		{nums: []int{2, 1, -1}, index: 0},
	} {
		index := pivotIndex(tc.nums)
		assert.Equal(t, tc.index, index)
	}
}
