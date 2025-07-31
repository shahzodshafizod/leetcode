package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		min  int
	}{
		{nums: []int{2, 1, 3, 4}, k: 1, min: 2},
		{nums: []int{2, 0, 2, 0}, k: 0, min: 0},
	} {
		minimum := minOperations(tc.nums, tc.k)
		assert.Equal(t, tc.min, minimum)
	}
}
