package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestValidMountainArray$
func TestValidMountainArray(t *testing.T) {
	for _, tc := range []struct {
		arr   []int
		valid bool
	}{
		{arr: []int{2, 1}, valid: false},
		{arr: []int{3, 5, 5}, valid: false},
		{arr: []int{0, 3, 2, 1}, valid: true},
		{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, valid: false},
		{arr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, valid: false},
	} {
		valid := validMountainArray(tc.arr)
		assert.Equal(t, tc.valid, valid)
	}
}
