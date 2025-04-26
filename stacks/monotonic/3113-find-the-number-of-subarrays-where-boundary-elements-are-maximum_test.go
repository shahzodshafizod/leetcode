package monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/monotonic/ -run ^TestNumberOfSubarrays$
func TestNumberOfSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int64
	}{
		{nums: []int{1, 4, 3, 3, 2}, count: 6},
		{nums: []int{3, 3, 3}, count: 6},
		{nums: []int{1}, count: 1},
	} {
		count := numberOfSubarrays(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
