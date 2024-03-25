package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFirstMissingPositive$
func TestFirstMissingPositive(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		missing int
	}{
		{nums: []int{1, 2, 0}, missing: 3},
		{nums: []int{3, 4, -1, 1}, missing: 2},
		{nums: []int{7, 8, 9, 11, 12}, missing: 1},
		{nums: []int{1}, missing: 2},
		{nums: []int{1, 2, 3, 4, 5}, missing: 6},
		{nums: []int{3, -3, 6, 3}, missing: 1},
		{nums: []int{1, 1}, missing: 2},
	} {
		missing := firstMissingPositive(tc.nums)
		assert.Equal(t, tc.missing, missing)
	}
}
