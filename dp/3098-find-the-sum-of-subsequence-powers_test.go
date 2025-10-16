package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSumOfPowers$
func TestSumOfPowers(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		total int
	}{
		{nums: []int{1, 2, 3, 4}, k: 3, total: 4},
		{nums: []int{2, 2}, k: 2, total: 0},
		{nums: []int{4, 3, -1}, k: 2, total: 10},
		// {nums: []int{-24, -921, 119, -291, -65, -628, 372, 274, 962, -592, -10, 67, -977, 85, -294, 349, -119, -846, -959, -79, -877, 833, 857, 44, 826, -295, -855, 554, -999, 759, -653, -423, -599, -928}, k: 19, total: 990202285},
	} {
		total := sumOfPowers(tc.nums, tc.k)
		assert.Equal(t, tc.total, total)
	}
}
