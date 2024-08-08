package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCanBeEqual$
func TestCanBeEqual(t *testing.T) {
	for _, tc := range []struct {
		target []int
		arr    []int
		canB   bool
	}{
		{target: []int{1, 2, 3, 4}, arr: []int{2, 4, 1, 3}, canB: true},
		{target: []int{7}, arr: []int{7}, canB: true},
		{target: []int{3, 7, 9}, arr: []int{3, 7, 11}, canB: false},
	} {
		canB := canBeEqual(tc.target, tc.arr)
		assert.Equal(t, tc.canB, canB)
	}
}
