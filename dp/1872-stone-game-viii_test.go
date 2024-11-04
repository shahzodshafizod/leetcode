package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestStoneGameVIII$
func TestStoneGameVIII(t *testing.T) {
	for _, tc := range []struct {
		stones     []int
		difference int
	}{
		{stones: []int{-10, -12}, difference: -22},
		{stones: []int{-1, 2, -3, 4, -5}, difference: 5},
		{stones: []int{7, -6, 5, 10, 5, -2, -6}, difference: 13},
		{stones: []int{25, -35, -37, 4, 34, 43, 16, -33, 0, -17, -31, -42, -42, 38, 12, -5, -43, -10, -37, 12}, difference: 38},
	} {
		difference := stoneGameVIII(tc.stones)
		assert.Equal(t, tc.difference, difference)
	}
}
