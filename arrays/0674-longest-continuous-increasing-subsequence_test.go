package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindLengthOfLCIS$
func TestFindLengthOfLCIS(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		result int
	}{
		{nums: []int{1, 3, 5, 4, 7}, result: 3},
		{nums: []int{2, 2, 2, 2, 2}, result: 1},
		{nums: []int{1}, result: 1},
		{nums: []int{1, 2, 3, 4, 5}, result: 5},
		{nums: []int{5, 4, 3, 2, 1}, result: 1},
		{nums: []int{1, 3, 5, 4, 7, 8, 9}, result: 4},
	} {
		result := findLengthOfLCIS(tc.nums)
		assert.Equal(t, tc.result, result)
	}
}
