package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestCountSubarrays$
func TestCountSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		minK  int
		maxK  int
		count int64
	}{
		{nums: []int{1, 1, 1, 1}, minK: 1, maxK: 1, count: 10},
		{nums: []int{1, 3, 5, 2, 7, 5}, minK: 1, maxK: 5, count: 2},
		{nums: []int{1, 2, 1, 4, 1, 3, 4}, minK: 1, maxK: 4, count: 16},
		{nums: []int{1, 3, 2, 2, 1, 3, 2, 2}, minK: 1, maxK: 3, count: 20},
	} {
		count := countSubarrays(tc.nums, tc.minK, tc.maxK)
		assert.Equal(t, tc.count, count)
	}
}
