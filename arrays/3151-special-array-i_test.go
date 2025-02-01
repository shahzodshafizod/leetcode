package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestIsArraySpecial$
func TestIsArraySpecial(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		special bool
	}{
		{nums: []int{1}, special: true},
		{nums: []int{2, 1, 4}, special: true},
		{nums: []int{4, 3, 1, 6}, special: false},
		{nums: []int{2}, special: true},
		{nums: []int{3}, special: true},
		{nums: []int{1, 5, 7}, special: false},
		{nums: []int{4, 8, 6, 2}, special: false},
		{nums: []int{6, 4, 8, 8, 2, 4, 2, 4, 8, 2, 2, 4, 4, 6, 5}, special: false},
	} {
		special := isArraySpecial(tc.nums)
		assert.Equal(t, tc.special, special)
	}
}
