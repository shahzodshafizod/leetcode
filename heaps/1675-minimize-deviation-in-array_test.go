package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestMinimumDeviation$
func TestMinimumDeviation(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		deviation int
	}{
		{nums: []int{1, 2, 3, 4}, deviation: 1},
		{nums: []int{4, 1, 5, 20, 3}, deviation: 3},
		{nums: []int{2, 10, 8}, deviation: 3},
		{nums: []int{2, 2, 3, 2}, deviation: 1},
		{nums: []int{3, 5}, deviation: 1},
	} {
		deviation := minimumDeviation(tc.nums)
		assert.Equal(t, tc.deviation, deviation)
	}
}
