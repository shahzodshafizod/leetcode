package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestWaysToSplitArray$
func TestWaysToSplitArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		ways int
	}{
		{nums: []int{10, 4, -8, 7}, ways: 2},
		{nums: []int{2, 3, 1, 0}, ways: 2},
		{nums: []int{-100000, 100000}, ways: 0},
	} {
		ways := waysToSplitArray(tc.nums)
		assert.Equal(t, tc.ways, ways)
	}
}
