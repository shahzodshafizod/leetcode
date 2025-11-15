package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMajorityElement$
func TestMajorityElement(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		mode int
	}{
		{nums: []int{3, 2, 3}, mode: 3},
		{nums: []int{2, 2, 1, 1, 1, 2, 2}, mode: 2},
	} {
		mode := majorityElement(tc.nums)
		assert.Equal(t, tc.mode, mode)
	}
}
