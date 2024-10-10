package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSplitArraySameAverage$
func TestSplitArraySameAverage(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		possible bool
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, possible: true},
		{nums: []int{3, 1}, possible: false},
		{nums: []int{17, 5, 5, 1, 14, 10, 13, 1, 6}, possible: true},
		{nums: []int{3, 4, 9, 4, 4, 3, 9, 8, 5, 3}, possible: true},
		{nums: []int{5, 16, 4, 11, 4}, possible: true},
		{nums: []int{2, 12, 18, 16, 19, 3}, possible: false},
		{nums: []int{5, 3, 11, 19, 2}, possible: true},
		{nums: []int{1, 6, 1}, possible: false},
		{nums: []int{1}, possible: false},
		{nums: []int{2}, possible: false},
		{nums: []int{0, 0, 0, 0, 0}, possible: true},
		{nums: []int{0}, possible: false},
		// {nums: []int{3322, 959, 598, 506, 4442, 3594, 8382, 7777, 7002, 7078, 2278, 4902, 1276, 1}, possible: false},
		// {nums: []int{904, 8738, 6439, 1889, 138, 5771, 8899, 5790, 662, 8402, 3074, 1844, 5926, 8720, 7159, 6793, 7402, 9466, 1282, 1748, 434, 842, 22}, possible: false},
		// {nums: []int{60, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}, possible: false},
		// {nums: []int{6795, 3310, 8624, 475, 7609, 7858, 7086, 8934, 6197, 2431, 3310, 760, 1432, 7518, 7068, 7182, 2681, 2679, 6461, 9928, 9651, 3258, 9346, 1666, 5400, 8384, 2751, 1234, 2183, 3520}, possible: false},
	} {
		possible := splitArraySameAverage(tc.nums)
		assert.Equal(t, tc.possible, possible)
	}
}
