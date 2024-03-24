package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindDuplicate$
func TestFindDuplicate(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		duplicate int
	}{
		{nums: []int{1, 3, 4, 2, 2}, duplicate: 2},
		{nums: []int{3, 1, 3, 4, 2}, duplicate: 3},
		{nums: []int{3, 3, 3, 3, 3}, duplicate: 3},
		{nums: []int{2, 2, 2, 2, 2}, duplicate: 2},
		{nums: []int{1, 1}, duplicate: 1},
	} {
		duplicate := findDuplicate(tc.nums)
		assert.Equal(t, tc.duplicate, duplicate)
	}
}
