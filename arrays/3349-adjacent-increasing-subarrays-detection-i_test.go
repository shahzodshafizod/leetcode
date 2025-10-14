package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestHasIncreasingSubarrays$
func TestHasIncreasingSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		has  bool
	}{
		{nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, k: 3, has: true},
		{nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}, k: 5, has: false},
	} {
		has := hasIncreasingSubarrays(tc.nums, tc.k)
		assert.Equal(t, tc.has, has)
	}
}
