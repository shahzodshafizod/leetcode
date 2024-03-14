package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestNumSubarraysWithSum$
func TestNumSubarraysWithSum(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		goal  int
		count int
	}{
		{nums: []int{1, 0, 1, 0, 1}, goal: 2, count: 4},
		{nums: []int{0, 0, 0, 0, 0}, goal: 0, count: 15},
	} {
		count := numSubarraysWithSum(tc.nums, tc.goal)
		assert.Equal(t, tc.count, count)
	}
}
