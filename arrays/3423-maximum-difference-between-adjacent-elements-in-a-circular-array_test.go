package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMaxAdjacentDistance$
func TestMaxAdjacentDistance(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		diff int
	}{
		{nums: []int{1, 2, 4}, diff: 3},
		{nums: []int{-5, -10, -5}, diff: 5},
	} {
		diff := maxAdjacentDistance(tc.nums)
		assert.Equal(t, tc.diff, diff)
	}
}
