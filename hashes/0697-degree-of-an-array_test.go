package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindShortestSubArray$
func TestFindShortestSubArray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 3, 2, 2, 3, 1}, 2},
	} {
		output := findShortestSubArray(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
