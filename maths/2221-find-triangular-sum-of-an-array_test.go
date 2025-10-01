package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestTriangularSum$
func TestTriangularSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		tsum int
	}{
		{nums: []int{1, 2, 3, 4, 5}, tsum: 8},
		{nums: []int{5}, tsum: 5},
	} {
		tsum := triangularSum(tc.nums)
		assert.Equal(t, tc.tsum, tsum)
	}
}
