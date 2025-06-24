package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestFindKDistantIndices$
func TestFindKDistantIndices(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		key     int
		k       int
		indices []int
	}{
		{nums: []int{3, 4, 9, 1, 3, 9, 5}, key: 9, k: 1, indices: []int{1, 2, 3, 4, 5, 6}},
		{nums: []int{2, 2, 2, 2, 2}, key: 2, k: 2, indices: []int{0, 1, 2, 3, 4}},
	} {
		indices := findKDistantIndices(tc.nums, tc.key, tc.k)
		assert.Equal(t, tc.indices, indices)
	}
}
