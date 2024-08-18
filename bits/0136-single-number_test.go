package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestSingleNumber$
func TestSingleNumber(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		unique int
	}{
		{nums: []int{2, 2, 1}, unique: 1},
		{nums: []int{4, 1, 2, 1, 2}, unique: 4},
		{nums: []int{1}, unique: 1},
	} {
		unique := singleNumber(tc.nums)
		assert.Equal(t, tc.unique, unique)
	}
}
