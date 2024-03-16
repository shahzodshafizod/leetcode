package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindMaxLength$
func TestFindMaxLength(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		length int
	}{
		{nums: []int{0, 1}, length: 2},
		{nums: []int{0, 1, 0}, length: 2},
		{nums: []int{0, 0, 0, 1, 1, 1, 1}, length: 6},
		{nums: []int{0, 1, 0, 1}, length: 4},
		{nums: []int{0, 1, 0, 1, 1}, length: 4},
		{nums: []int{0, 0, 1, 0, 0, 0, 1, 1}, length: 6},
		{nums: []int{1, 0, 0, 1, 1, 1, 0, 1, 0, 0}, length: 10},
		{nums: []int{0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0}, length: 6},
	} {
		length := findMaxLength(tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
