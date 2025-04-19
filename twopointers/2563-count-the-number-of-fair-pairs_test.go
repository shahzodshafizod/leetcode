package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestCountFairPairs$
func TestCountFairPairs(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		lower int
		upper int
		pairs int64
	}{
		{nums: []int{0, 1, 7, 4, 4, 5}, lower: 3, upper: 6, pairs: 6},
		{nums: []int{1, 7, 9, 2, 5}, lower: 11, upper: 11, pairs: 1},
	} {
		pairs := countFairPairs(tc.nums, tc.lower, tc.upper)
		assert.Equal(t, tc.pairs, pairs)
	}
}
