package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestLargestDivisibleSubset$
func TestLargestDivisibleSubset(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		answer []int
	}{
		{nums: []int{1, 2, 3}, answer: []int{1, 2}},
		{nums: []int{1, 2, 4, 8}, answer: []int{1, 2, 4, 8}},
	} {
		answer := largestDivisibleSubset(tc.nums)
		assert.Equal(t, tc.answer, answer)
	}
}
