package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestProductExceptSelf$
func TestProductExceptSelf(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		answer []int
	}{
		{nums: []int{1, 2, 3, 4}, answer: []int{24, 12, 8, 6}},
		{nums: []int{-1, 1, 0, -3, 3}, answer: []int{0, 0, 9, 0, 0}},
	} {
		answer := productExceptSelf(tc.nums)
		assert.Equal(t, tc.answer, answer)
	}
}
