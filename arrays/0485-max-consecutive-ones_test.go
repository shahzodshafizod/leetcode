package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindMaxConsecutiveOnes$
func TestFindMaxConsecutiveOnes(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 1, 0, 1, 1, 1}, output: 3},
		{nums: []int{1, 0, 1, 1, 0, 1}, output: 2},
		{nums: []int{0, 0, 0}, output: 0},
		{nums: []int{1, 1, 1, 1}, output: 4},
	} {
		output := findMaxConsecutiveOnes(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
