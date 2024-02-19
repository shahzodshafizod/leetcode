package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestLastStoneWeight$
func TestLastStoneWeight(t *testing.T) {
	for _, tc := range []struct {
		stones []int
		weight int
	}{
		{stones: []int{2, 7, 4, 1, 8, 1}, weight: 1},
		{stones: []int{1}, weight: 1},
		{stones: []int{8, 7, 1}, weight: 0},
	} {
		weight := lastStoneWeight(tc.stones)
		assert.Equal(t, tc.weight, weight)
	}
}
