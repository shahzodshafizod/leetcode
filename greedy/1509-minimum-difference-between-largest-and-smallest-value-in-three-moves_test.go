package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinDifference$
func TestMinDifference(t *testing.T) {
	for _, tc := range []struct {
		nums       []int
		difference int
	}{
		{nums: []int{5, 3, 2, 4}, difference: 0},
		{nums: []int{1, 5, 0, 10, 14}, difference: 1},
		{nums: []int{3, 100, 20}, difference: 0},
	} {
		difference := minDifference(tc.nums)
		assert.Equal(t, tc.difference, difference)
	}
}
