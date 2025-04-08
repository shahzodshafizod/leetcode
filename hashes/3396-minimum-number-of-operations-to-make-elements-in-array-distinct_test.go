package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMinimumOperations$
func TestMinimumOperations(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		opers int
	}{
		{nums: []int{1, 2, 3, 4, 2, 3, 3, 5, 7}, opers: 2},
		{nums: []int{4, 5, 6, 4, 4}, opers: 2},
		{nums: []int{6, 7, 8, 9}, opers: 0},
	} {
		opers := minimumOperations(tc.nums)
		assert.Equal(t, tc.opers, opers)
	}
}
