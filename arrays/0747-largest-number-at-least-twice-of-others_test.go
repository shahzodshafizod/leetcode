package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestDominantIndex$
func TestDominantIndex(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1}, 0},
		{[]int{0, 0, 0, 1}, 3},
		{[]int{1, 0}, 0},
		{[]int{0, 0, 3, 2}, -1},
	} {
		output := dominantIndex(tc.nums)
		assert.Equal(t, tc.expected, output)
	}
}
