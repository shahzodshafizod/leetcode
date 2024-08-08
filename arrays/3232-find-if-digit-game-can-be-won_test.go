package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCanAliceWin$
func TestCanAliceWin(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		canWin bool
	}{
		{nums: []int{1, 2, 3, 4, 10}, canWin: false},
		{nums: []int{1, 2, 3, 4, 5, 14}, canWin: true},
		{nums: []int{5, 5, 5, 25}, canWin: true},
	} {
		canWin := canAliceWin(tc.nums)
		assert.Equal(t, tc.canWin, canWin)
	}
}
