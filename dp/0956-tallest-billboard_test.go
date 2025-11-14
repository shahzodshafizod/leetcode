package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestTallestBillboard$
func TestTallestBillboard(t *testing.T) {
	for _, tc := range []struct {
		rods   []int
		height int
	}{
		{rods: []int{1, 2, 3, 6}, height: 6},
		{rods: []int{1, 2, 3, 4, 5, 6}, height: 10},
		{rods: []int{1, 2}, height: 0},
		{rods: []int{96, 112, 101, 100, 104, 93, 106, 99, 114, 81, 94}, height: 503},
		{rods: []int{142, 178, 178, 143, 133, 139, 117, 153, 144, 162, 160, 147, 136, 149, 163, 160, 130, 157, 159}, height: 1425},
		{rods: []int{33, 30, 41, 34, 37, 29, 26, 31, 42, 33, 38, 27, 33, 31, 35, 900, 900, 900, 900, 900}, height: 2050},
	} {
		height := tallestBillboard(tc.rods)
		assert.Equal(t, tc.height, height)
	}
}
