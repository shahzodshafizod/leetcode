package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSubsets$
func TestSubsets(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		powersets [][]int
	}{
		// iteratively
		{
			nums:      []int{1, 2, 3},
			powersets: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			nums:      []int{0},
			powersets: [][]int{{}, {0}},
		},
		{
			nums:      []int{9, 0, 3, 5, 7},
			powersets: [][]int{{}, {9}, {0}, {9, 0}, {3}, {9, 3}, {0, 3}, {9, 0, 3}, {5}, {9, 5}, {0, 5}, {9, 0, 5}, {3, 5}, {9, 3, 5}, {0, 3, 5}, {9, 0, 3, 5}, {7}, {9, 7}, {0, 7}, {9, 0, 7}, {3, 7}, {9, 3, 7}, {0, 3, 7}, {9, 0, 3, 7}, {5, 7}, {9, 5, 7}, {0, 5, 7}, {9, 0, 5, 7}, {3, 5, 7}, {9, 3, 5, 7}, {0, 3, 5, 7}, {9, 0, 3, 5, 7}},
		},

		// // recursively
		// {
		// 	nums:      []int{1, 2, 3},
		// 	powersets: [][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}},
		// },
		// {
		// 	nums:      []int{0},
		// 	powersets: [][]int{{0}, {}},
		// },
	} {
		powersets := subsets(tc.nums)
		assert.Equal(t, tc.powersets, powersets)
	}
}
