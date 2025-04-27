package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountSubarrays$
func TestCountSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{1, 2, 1, 4, 1}, count: 1},
		{nums: []int{1, 1, 1}, count: 0},
	} {
		count := countSubarrays(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
