package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		output int
	}{
		{nums: []int{1, 2, 3, 4}, k: 5, output: 0},
		{nums: []int{1, 2, 3}, k: 7, output: 6},
		{nums: []int{5, 6, 7}, k: 9, output: 0},
		{nums: []int{1}, k: 3, output: 1},
	} {
		output := minOperations(tc.nums, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
