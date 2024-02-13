package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestShuffle$
func TestShuffle(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		n        int
		shuffled []int
	}{
		{nums: []int{2, 5, 1, 3, 4, 7}, n: 3, shuffled: []int{2, 3, 5, 4, 1, 7}},
		{nums: []int{1, 2, 3, 4, 4, 3, 2, 1}, n: 4, shuffled: []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, n: 4, shuffled: []int{1, 5, 2, 6, 3, 7, 4, 8}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, n: 5, shuffled: []int{1, 6, 2, 7, 3, 8, 4, 9, 5, 0}},
		{nums: []int{1, 1, 2, 2}, n: 2, shuffled: []int{1, 2, 1, 2}},
	} {
		shuffled := shuffle(tc.nums, tc.n)
		assert.Equal(t, tc.shuffled, shuffled)
	}
}
