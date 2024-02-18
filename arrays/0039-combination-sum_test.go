package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCombinationSum$
func TestCombinationSum(t *testing.T) {
	for _, tc := range []struct {
		candidates   []int
		target       int
		combinations [][]int
	}{
		{
			candidates:   []int{2, 3, 6, 7},
			target:       7,
			combinations: [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates:   []int{2, 3, 5},
			target:       8,
			combinations: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates:   []int{2},
			target:       1,
			combinations: [][]int{},
		},
		{
			candidates:   []int{2, 3},
			target:       6,
			combinations: [][]int{{2, 2, 2}, {3, 3}},
		},
	} {
		combinations := combinationSum(tc.candidates, tc.target)
		assert.Equal(t, tc.combinations, combinations)
	}
}
