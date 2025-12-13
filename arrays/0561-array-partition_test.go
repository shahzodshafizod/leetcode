package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestArrayPairSum$
func TestArrayPairSum(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 4, 3, 2}, output: 4},
		{nums: []int{6, 2, 6, 5, 1, 2}, output: 9},
		{nums: []int{1, 2}, output: 1},
	} {
		output := arrayPairSum(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
