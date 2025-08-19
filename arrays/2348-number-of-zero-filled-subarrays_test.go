package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestZeroFilledSubarray$
func TestZeroFilledSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int64
	}{
		{nums: []int{1, 3, 0, 0, 2, 0, 0, 4}, count: 6},
		{nums: []int{0, 0, 0, 2, 0, 0}, count: 9},
		{nums: []int{2, 10, 2019}, count: 0},
	} {
		count := zeroFilledSubarray(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
