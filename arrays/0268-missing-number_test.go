package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMissingNumber$
func TestMissingNumber(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		missing int
	}{
		{nums: []int{3, 0, 1}, missing: 2},
		{nums: []int{0, 1}, missing: 2},
		{nums: []int{2, 1}, missing: 0},
		{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, missing: 8},
	} {
		missing := missingNumber(tc.nums)
		assert.Equal(t, tc.missing, missing)
	}
}
