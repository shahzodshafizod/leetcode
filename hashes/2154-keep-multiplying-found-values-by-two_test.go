package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindFinalValue$
func TestFindFinalValue(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		original int
		output   int
	}{
		{nums: []int{5, 3, 6, 1, 12}, original: 3, output: 24},
		{nums: []int{2, 7, 9}, original: 4, output: 4},
		{nums: []int{8, 19, 4, 2, 15, 3}, original: 2, output: 16},
		{nums: []int{1}, original: 1, output: 2},
		{nums: []int{1, 2}, original: 1, output: 4},
	} {
		output := findFinalValue(tc.nums, tc.original)
		assert.Equal(t, tc.output, output)
	}
}
