package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaximumLength$
func TestMaximumLength(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		length int
	}{
		{nums: []int{1, 2, 3, 4}, length: 4},
		{nums: []int{1, 2, 1, 1, 2, 1, 2}, length: 6},
		{nums: []int{1, 3}, length: 2},
	} {
		length := maximumLength(tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
