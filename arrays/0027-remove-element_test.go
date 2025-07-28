package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestRemoveElement$
func TestRemoveElement(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range []struct {
		nums         []int
		val          int
		expectedNums []int
		k            int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, expectedNums: []int{2, 2, 2, 3}, k: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expectedNums: []int{0, 1, 3, 0, 4, 2, 4, 2}, k: 5},
		{nums: []int{0, 2, 3, 4, 5, 3, 1, 0}, val: 3, expectedNums: []int{0, 2, 4, 5, 1, 0, 0, 0}, k: 6},
		{nums: []int{3, 3}, val: 5, expectedNums: []int{3, 3}, k: 2},
		// two pointers
		// {nums: []int{3, 2, 2, 3}, val: 3, expectedNums: []int{2, 2, 3, 3}, k: 2},
		// {nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expectedNums: []int{0, 1, 4, 0, 3, 2, 2, 2}, k: 5},
		// {nums: []int{0, 2, 3, 4, 5, 3, 1, 0}, val: 3, expectedNums: []int{0, 2, 0, 4, 5, 1, 3, 3}, k: 6},
	} {
		k := removeElement(tc.nums, tc.val)
		assert.Equal(tc.k, k)

		for i := 0; i < k; i++ {
			assert.Equal(tc.expectedNums[i], tc.nums[i])
		}
	}
}
