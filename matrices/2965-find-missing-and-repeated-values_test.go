package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestFindMissingAndRepeatedValues$
func TestFindMissingAndRepeatedValues(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		ans  []int
	}{
		{grid: [][]int{{1, 3}, {2, 2}}, ans: []int{2, 4}},
		{grid: [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}, ans: []int{9, 5}},
	} {
		ans := findMissingAndRepeatedValues(tc.grid)
		assert.Equal(t, tc.ans, ans)
	}
}
