package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountMatchingSubarrays$
func TestCountMatchingSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		pattern []int
		output  int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6}, pattern: []int{1, 1}, output: 4},
		{nums: []int{1, 4, 4, 1, 3, 5, 5, 3}, pattern: []int{1, 0, -1}, output: 2},
		{nums: []int{1, 2, 3}, pattern: []int{1}, output: 2},
		{nums: []int{5, 5, 5, 5}, pattern: []int{0, 0}, output: 2},
	} {
		output := countMatchingSubarrays(tc.nums, tc.pattern)
		assert.Equal(t, tc.output, output)
	}
}
