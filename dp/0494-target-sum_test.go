package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestFindTargetSumWays$
func TestFindTargetSumWays(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		ways   int
	}{
		{nums: []int{1}, target: 1, ways: 1},
		{nums: []int{1}, target: 2, ways: 0},
		{nums: []int{100, 100}, target: -300, ways: 0},
		{nums: []int{1, 1, 1, 1, 1}, target: 3, ways: 5},
		// {nums: []int{12, 25, 42, 49, 41, 15, 22, 34, 28, 31}, target: 35, ways: 8},
		// {nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, target: 1, ways: 524288},
		// {nums: []int{3, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 2, 7, 9, 13, 27, 31, 37, 47, 53}, target: 107, ways: 0},
		// {nums: []int{0, 5, 22, 39, 30, 5, 40, 23, 47, 43, 19, 36, 10, 28, 46, 14, 11, 5, 50, 41}, target: 48, ways: 5726},
		// {nums: []int{30, 29, 48, 14, 29, 4, 31, 12, 40, 13, 30, 1, 5, 32, 16, 17, 13, 20, 41, 38}, target: 9, ways: 6867},
		// {nums: []int{3, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 2, 107, 109, 113, 127, 131, 137, 47, 53}, target: 4, ways: 2780},
	} {
		ways := findTargetSumWays(tc.nums, tc.target)
		assert.Equal(t, tc.ways, ways)
	}
}
