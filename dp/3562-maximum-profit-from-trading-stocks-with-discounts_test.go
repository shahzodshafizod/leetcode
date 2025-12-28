package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaximumProfit$
func TestMaximumProfit(t *testing.T) {
	for _, tc := range []struct {
		n         int
		present   []int
		future    []int
		hierarchy [][]int
		budget    int
		expected  int
	}{
		{n: 2, present: []int{1, 2}, future: []int{4, 3}, hierarchy: [][]int{{1, 2}}, budget: 3, expected: 5},
		{n: 2, present: []int{3, 4}, future: []int{5, 8}, hierarchy: [][]int{{1, 2}}, budget: 4, expected: 4},
		{n: 3, present: []int{4, 6, 8}, future: []int{7, 9, 11}, hierarchy: [][]int{{1, 2}, {1, 3}}, budget: 10, expected: 10},
		{n: 3, present: []int{5, 2, 3}, future: []int{8, 5, 6}, hierarchy: [][]int{{1, 2}, {2, 3}}, budget: 7, expected: 12},
		{n: 1, present: []int{45}, future: []int{38}, hierarchy: [][]int{}, budget: 115, expected: 0},
		{n: 2, present: []int{6, 11}, future: []int{5, 48}, hierarchy: [][]int{{1, 2}}, budget: 142, expected: 42},
		{n: 3, present: []int{42, 27, 32}, future: []int{46, 8, 17}, hierarchy: [][]int{{1, 2}, {2, 3}}, budget: 93, expected: 4},
	} {
		result := maximumProfit(tc.n, tc.present, tc.future, tc.hierarchy, tc.budget)
		assert.Equal(t, tc.expected, result)
	}
}
