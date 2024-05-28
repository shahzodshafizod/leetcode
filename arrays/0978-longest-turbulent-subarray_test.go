package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMaxTurbulenceSize$
func TestMaxTurbulenceSize(t *testing.T) {
	for _, tc := range []struct {
		arr []int
		max int
	}{
		{arr: []int{9, 4, 2, 10, 7, 8, 8, 1, 9}, max: 5},
		{arr: []int{9, 4, 2, 10, 7, 8}, max: 5},
		{arr: []int{4, 8, 12, 16}, max: 2},
		{arr: []int{100}, max: 1},
		{arr: []int{100, 100, 100}, max: 1},
		{arr: []int{100, 100, 100, 2}, max: 2},
	} {
		max := maxTurbulenceSize(tc.arr)
		assert.Equal(t, tc.max, max)
	}
}
