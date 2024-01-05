package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dsa/arrays/ -run ^TestMaxArea$
func TestMaxArea(t *testing.T) {
	for _, data := range []struct {
		height []int
		max    int
	}{
		{[]int{7, 1, 2, 3, 9}, 28},
		{[]int{}, 0},
		{[]int{7}, 0},
		{[]int{6, 9, 3, 4, 5, 8}, 32},
		{[]int{4, 8, 1, 2, 3, 9}, 32},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	} {
		max := maxArea(data.height)
		assert.Equal(t, data.max, max)
	}
}
