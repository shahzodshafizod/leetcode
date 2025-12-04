package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMinimumDifference$
func TestMinimumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		output int
	}{
		{nums: []int{1, 2, 4, 5}, k: 3, output: 0},
		{nums: []int{1, 3, 1, 3}, k: 2, output: 1},
		{nums: []int{1}, k: 10, output: 9},
		{nums: []int{1, 2, 1, 2}, k: 4, output: 1},
		{nums: []int{3, 6, 9}, k: 5, output: 1},
	} {
		output := minimumDifference(tc.nums, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
