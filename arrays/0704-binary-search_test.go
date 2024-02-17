package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSearch$
func TestSearch(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		index  int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, index: 4},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, index: -1},
	} {
		index := search(tc.nums, tc.target)
		assert.Equal(t, tc.index, index)
	}
}
