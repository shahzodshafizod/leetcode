package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMaxSlidingWindow$
func TestMaxSlidingWindow(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		window []int
	}{
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, window: []int{3, 3, 5, 5, 6, 7}},
		{nums: []int{1}, k: 1, window: []int{1}},
		{nums: []int{2, 1, 3, 4, 6, 3, 8, 9, 10, 12, 56}, k: 4, window: []int{4, 6, 6, 8, 9, 10, 12, 56}},
		{nums: []int{1, -1}, k: 1, window: []int{1, -1}},
		{nums: []int{9, 10, 9, -7, -4, -8, 2, -6}, k: 5, window: []int{10, 10, 9, 2}}, // tricky test case
		{nums: []int{7, 2, 4}, k: 2, window: []int{7, 4}},
		{nums: []int{-7, -8, 7, 5, 7, 1, 6, 0}, k: 4, window: []int{7, 7, 7, 7, 7}},
		// {nums: []int{}, k: 0, window: []int{}}, // invalid test case due to constrains: 1 <= k <= nums.length
	} {
		window := maxSlidingWindow(tc.nums, tc.k)
		assert.Equal(t, tc.window, window)
	}
}
