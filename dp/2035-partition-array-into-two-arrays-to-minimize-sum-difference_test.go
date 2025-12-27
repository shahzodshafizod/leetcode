package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinimumDifference$
func TestMinimumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{3, 9, 7, 3}, output: 2},
		{nums: []int{-36, 36}, output: 72},
		{nums: []int{2, -1, 0, 4, -2, -9}, output: 0},
	} {
		output := minimumDifference(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
