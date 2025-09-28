package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestLargestPerimeter$
func TestLargestPerimeter(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		perimeter int
	}{
		{nums: []int{2, 1, 2}, perimeter: 5},
		{nums: []int{1, 2, 1, 10}, perimeter: 0},
	} {
		perimeter := largestPerimeter(tc.nums)
		assert.Equal(t, tc.perimeter, perimeter)
	}
}
