package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMaxAscendingSum$
func TestMaxAscendingSum(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		maxsum int
	}{
		{nums: []int{10, 20, 30, 5, 10, 50}, maxsum: 65},
		{nums: []int{10, 20, 30, 40, 50}, maxsum: 150},
		{nums: []int{12, 17, 15, 13, 10, 11, 12}, maxsum: 33},
	} {
		maxsum := maxAscendingSum(tc.nums)
		assert.Equal(t, tc.maxsum, maxsum)
	}
}
