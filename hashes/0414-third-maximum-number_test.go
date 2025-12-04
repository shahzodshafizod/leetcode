package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestThirdMax$
func TestThirdMax(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		tmax int
	}{
		{nums: []int{3, 2, 1}, tmax: 1},
		{nums: []int{1, 2}, tmax: 2},
		{nums: []int{2, 2, 3, 1}, tmax: 1},
	} {
		tmax := thirdMax(tc.nums)
		assert.Equal(t, tc.tmax, tmax)
	}
}
