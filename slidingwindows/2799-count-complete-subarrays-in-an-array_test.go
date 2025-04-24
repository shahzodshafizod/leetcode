package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestCountCompleteSubarrays$
func TestCountCompleteSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{1, 3, 1, 2, 2}, count: 4},
		{nums: []int{5, 5, 5, 5}, count: 10},
		{nums: []int{1898, 370, 822, 1659, 1360, 128, 370, 360, 261, 1898}, count: 4},
	} {
		count := countCompleteSubarrays(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
