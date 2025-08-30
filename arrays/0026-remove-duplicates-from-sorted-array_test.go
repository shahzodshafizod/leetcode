package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestRemoveDuplicates$
func TestRemoveDuplicates(t *testing.T) {
	for _, tc := range []struct {
		nums         []int
		expectedNums []int
		k            int
	}{
		{nums: []int{1, 1, 2}, expectedNums: []int{1, 2, 0}, k: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expectedNums: []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0}, k: 5},
	} {
		k := removeDuplicates(tc.nums)
		assert.Equal(t, tc.k, k)

		for i := range k {
			assert.Equal(t, tc.expectedNums[i], tc.nums[i])
		}
	}
}
