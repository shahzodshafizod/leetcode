package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestMaximumProduct$
func TestMaximumProduct(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 2, 3}, output: 6},
		{nums: []int{1, 2, 3, 4}, output: 24},
		{nums: []int{-1, -2, -3}, output: -6},
		{nums: []int{-4, -3, -2, -1, 60}, output: 720}, // -4 * -3 * 60 = 720
		{nums: []int{1, 2, 3, 4, 5, 6}, output: 120},
		{nums: []int{-10, -10, 1, 3, 2}, output: 300}, // -10 * -10 * 3 = 300
	} {
		output := maximumProduct(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
