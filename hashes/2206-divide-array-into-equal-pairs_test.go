package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestDivideArray$
func TestDivideArray(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		divisable bool
	}{
		{nums: []int{3, 2, 3, 2, 2, 2}, divisable: true},
		{nums: []int{1, 2, 3, 4}, divisable: false},
		{nums: []int{1, 2, 4, 7}, divisable: false},
	} {
		divisable := divideArray(tc.nums)
		assert.Equal(t, tc.divisable, divisable)
	}
}
