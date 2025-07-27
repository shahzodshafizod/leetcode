package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountHillValley$
func TestCountHillValley(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{2, 4, 1, 1, 6, 5}, count: 3},
		{nums: []int{6, 6, 5, 5, 4, 1}, count: 0},
	} {
		count := countHillValley(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
