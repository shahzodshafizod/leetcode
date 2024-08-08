package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountSmaller$
func TestCountSmaller(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		smallers []int
	}{
		{nums: []int{5, 2, 6, 1}, smallers: []int{2, 1, 1, 0}},
		{nums: []int{-1}, smallers: []int{0}},
		{nums: []int{-1, -1}, smallers: []int{0, 0}},
		{nums: []int{9, 11, 15}, smallers: []int{0, 0, 0}},
		{nums: []int{15, 11, 9}, smallers: []int{2, 1, 0}},
		{nums: []int{2, 0, 1}, smallers: []int{2, 0, 0}},
	} {
		smallers := countSmaller(tc.nums)
		assert.Equal(t, tc.smallers, smallers)
	}
}
