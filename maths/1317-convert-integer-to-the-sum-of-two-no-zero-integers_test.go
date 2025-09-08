package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestGetNoZeroIntegers$
func TestGetNoZeroIntegers(t *testing.T) {
	for _, tc := range []struct {
		n    int
		nums []int
	}{
		{n: 2, nums: []int{1, 1}},
		{n: 11, nums: []int{5, 6}},
	} {
		nums := getNoZeroIntegers(tc.n)
		assert.Equal(t, tc.nums, nums)
	}
}
