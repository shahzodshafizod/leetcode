package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestSubsetXORSum$
func TestSubsetXORSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		sum  int
	}{
		{nums: []int{1, 3}, sum: 6},
		{nums: []int{5, 1, 6}, sum: 28},
		{nums: []int{3, 4, 5, 6, 7, 8}, sum: 480},
	} {
		sum := subsetXORSum(tc.nums)
		assert.Equal(t, tc.sum, sum)
	}
}
