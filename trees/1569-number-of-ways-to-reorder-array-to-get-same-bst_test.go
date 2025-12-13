package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestNumOfWays$
func TestNumOfWays(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{2, 1, 3}, output: 1},
		{nums: []int{3, 4, 5, 1, 2}, output: 5},
		{nums: []int{1, 2, 3}, output: 0},
		{nums: []int{3, 1, 2, 5, 4, 6}, output: 19},
		{nums: []int{9, 4, 2, 1, 3, 6, 5, 7, 8, 14, 11, 10, 12, 13, 16, 15, 17, 18}, output: 216212978},
	} {
		output := numOfWays(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
