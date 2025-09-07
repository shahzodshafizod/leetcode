package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSumZero$
func TestSumZero(t *testing.T) {
	for _, tc := range []struct {
		n    int
		nums []int
	}{
		{n: 5, nums: []int{0, 1, 2, -2, -1}},
		{n: 3, nums: []int{0, 1, -1}},
		{n: 1, nums: []int{0}},
		{n: 4, nums: []int{1, 2, -2, -1}},
	} {
		nums := sumZero(tc.n)
		assert.Equal(t, tc.nums, nums)
	}
}
