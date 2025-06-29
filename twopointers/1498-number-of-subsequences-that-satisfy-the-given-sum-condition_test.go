package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestNumSubseq$
func TestNumSubseq(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		count  int
	}{
		{nums: []int{3, 5, 6, 7}, target: 9, count: 4},
		{nums: []int{3, 3, 6, 8}, target: 10, count: 6},
		{nums: []int{2, 3, 3, 4, 6, 7}, target: 12, count: 61},
		{nums: []int{7, 10, 7, 3, 7, 5, 4}, target: 12, count: 56},
	} {
		count := numSubseq(tc.nums, tc.target)
		assert.Equal(t, tc.count, count)
	}
}
