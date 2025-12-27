package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestDistributeCandies$
func TestDistributeCandies(t *testing.T) {
	for _, tc := range []struct {
		candyType []int
		output    int
	}{
		{candyType: []int{1, 1, 2, 2, 3, 3}, output: 3},
		{candyType: []int{1, 1, 2, 3}, output: 2},
		{candyType: []int{6, 6, 6, 6}, output: 1},
		{candyType: []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3}, output: 3},
		{candyType: []int{100000, -100000, 100000, -100000}, output: 2},
	} {
		output := distributeCandies(tc.candyType)
		assert.Equal(t, tc.output, output)
	}
}
