package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindErrorNums$
func TestFindErrorNums(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{1, 2, 2, 4}, expected: []int{2, 3}},
		{nums: []int{1, 1}, expected: []int{1, 2}},
		{nums: []int{3, 2, 2}, expected: []int{2, 1}},
		{nums: []int{2, 2}, expected: []int{2, 1}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10}, expected: []int{10, 11}},
		{nums: []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9}, expected: []int{2, 10}},
	} {
		output := findErrorNums(tc.nums)
		assert.Equal(t, tc.expected, output)
	}
}
