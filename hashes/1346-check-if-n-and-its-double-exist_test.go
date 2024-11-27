package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCheckIfExist$
func TestCheckIfExist(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		exists bool
	}{
		{arr: []int{10, 2, 5, 3}, exists: true},
		{arr: []int{7, 1, 14, 11}, exists: true},
		{arr: []int{3, 1, 7, 11}, exists: false},
		{arr: []int{-2, 0, 10, -19, 4, 6, -8}, exists: false},
		{arr: []int{0}, exists: false},
		{arr: []int{0, 0}, exists: true},
	} {
		exists := checkIfExist(tc.arr)
		assert.Equal(t, tc.exists, exists)
	}
}
