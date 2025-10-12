package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMagicalSum$
func TestMagicalSum(t *testing.T) {
	for _, tc := range []struct {
		m    int
		k    int
		nums []int
		sum  int
	}{
		{m: 5, k: 5, nums: []int{1, 10, 100, 10000, 1000000}, sum: 991600007},
		{m: 2, k: 2, nums: []int{5, 4, 3, 2, 1}, sum: 170},
		{m: 1, k: 1, nums: []int{28}, sum: 28},
	} {
		sum := magicalSum(tc.m, tc.k, tc.nums)
		assert.Equal(t, tc.sum, sum)
	}
}
