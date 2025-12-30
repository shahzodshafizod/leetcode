package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinimumBoxes$
func TestMinimumBoxes(t *testing.T) {
	for _, tc := range []struct {
		apple    []int
		capacity []int
		output   int
	}{
		{[]int{1, 3, 2}, []int{4, 3, 1, 5, 2}, 2},
		{[]int{5, 5, 5}, []int{2, 4, 2, 7}, 4},
		{[]int{1}, []int{1}, 1},
		{[]int{10}, []int{5, 6}, 2},
		{[]int{1, 1, 1, 1, 1}, []int{10}, 1},
	} {
		output := minimumBoxes(tc.apple, tc.capacity)
		assert.Equal(t, tc.output, output)
	}
}
