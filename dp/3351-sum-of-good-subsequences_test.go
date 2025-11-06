package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSumOfGoodSubsequences$
func TestSumOfGoodSubsequences(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		sums int
	}{
		{nums: []int{1, 2, 1}, sums: 14},
		{nums: []int{3, 4, 5}, sums: 40},
	} {
		sums := sumOfGoodSubsequences(tc.nums)
		assert.Equal(t, tc.sums, sums)
	}
}
