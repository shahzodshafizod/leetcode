package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMinOperations3191$
func TestMinOperations3191(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		flips int
	}{
		{nums: []int{0, 1, 1, 1, 0, 0}, flips: 3},
		{nums: []int{0, 1, 1, 1}, flips: -1},
	} {
		flips := minOperations3191(tc.nums)
		assert.Equal(t, tc.flips, flips)
	}
}
