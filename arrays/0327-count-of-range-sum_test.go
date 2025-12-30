package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountRangeSum$
func TestCountRangeSum(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		lower  int
		upper  int
		output int
	}{
		{nums: []int{-2, 5, -1}, lower: -2, upper: 2, output: 3},
		{nums: []int{0}, lower: 0, upper: 0, output: 1},
		{nums: []int{2147483647, -2147483648, -1, 0}, lower: -1, upper: 0, output: 4},
		{nums: []int{0, -3, -3, 1, 1, 2}, lower: 3, upper: 5, output: 2},
		{nums: []int{-2, 5, -1}, lower: 0, upper: 2, output: 1},
	} {
		output := countRangeSum(tc.nums, tc.lower, tc.upper)
		assert.Equal(t, tc.output, output)
	}
}
