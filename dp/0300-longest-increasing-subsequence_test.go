package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestLengthOfLIS$
func TestLengthOfLIS(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		length int
	}{
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, length: 4},
		{nums: []int{0, 1, 0, 3, 2, 3}, length: 4},
		{nums: []int{7, 7, 7, 7, 7, 7, 7}, length: 1},
		{nums: []int{3, 1, 8, 2, 5}, length: 3},
		{nums: []int{5, 2, 8, 6, 3, 6, 9, 5}, length: 4},
		{nums: []int{4, 10, 4, 3, 8, 9}, length: 3},
		{nums: []int{0, 1, 0, 3, 2, 3}, length: 4},
		{nums: []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}, length: 6},
		{nums: []int{1, 4, 2, 3}, length: 3},
	} {
		length := lengthOfLIS(tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
