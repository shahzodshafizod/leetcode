package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSpecialArray$
func TestSpecialArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		x    int
	}{
		{nums: []int{3, 5}, x: 2},
		{nums: []int{3, 8, 14, 5}, x: -1},
		{nums: []int{0, 0}, x: -1},
		{nums: []int{0, 4, 3, 0, 4}, x: 3},
	} {
		x := specialArray(tc.nums)
		assert.Equal(t, tc.x, x)
	}
}
