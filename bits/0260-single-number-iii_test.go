package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestSingleNumberIII$
func TestSingleNumberIII(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		unique []int
	}{
		{nums: []int{1, 2, 1, 3, 2, 5}, unique: []int{3, 5}},
		{nums: []int{-1, 0}, unique: []int{-1, 0}},
		{nums: []int{0, 1}, unique: []int{1, 0}},
		{nums: []int{1, 1, 0, -2147483648}, unique: []int{-2147483648, 0}},
	} {
		unique := singleNumberIII(tc.nums)
		assert.Equal(t, tc.unique, unique)
	}
}
