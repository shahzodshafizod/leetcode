package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCheck$
func TestCheck(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		rotated bool
	}{
		{nums: []int{3, 4, 5, 1, 2}, rotated: true},
		{nums: []int{2, 1, 3, 4}, rotated: false},
		{nums: []int{1, 2, 3}, rotated: true},
		{nums: []int{1, 3, 2}, rotated: false},
		{nums: []int{2, 1}, rotated: true},
		{nums: []int{2, 4, 1, 3}, rotated: false},
		{nums: []int{2, 1, 2, 2, 2}, rotated: true},
		{nums: []int{6, 7, 2, 7, 5}, rotated: false},
		{nums: []int{7, 9, 1, 1, 1, 3}, rotated: true},
		{nums: []int{10, 10, 1, 4, 5, 10}, rotated: true},
	} {
		rotated := check(tc.nums)
		assert.Equal(t, tc.rotated, rotated)
	}
}
