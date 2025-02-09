package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCountBadPairs$
func TestCountBadPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		bads int64
	}{
		{nums: []int{4, 1, 3, 3}, bads: 5},
		{nums: []int{1, 2, 3, 4, 5}, bads: 0},
	} {
		bads := countBadPairs(tc.nums)
		assert.Equal(t, tc.bads, bads)
	}
}
