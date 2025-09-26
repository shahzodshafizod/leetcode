package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestTriangleNumber$
func TestTriangleNumber(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{2, 2, 3, 4}, count: 3},
		{nums: []int{4, 2, 3, 4}, count: 4},
	} {
		count := triangleNumber(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
