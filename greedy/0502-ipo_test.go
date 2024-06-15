package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestFindMaximizedCapital$
func TestFindMaximizedCapital(t *testing.T) {
	for _, tc := range []struct {
		k            int
		w            int
		profits      []int
		capital      []int
		finalCapital int
	}{
		{k: 2, w: 0, profits: []int{1, 2, 3}, capital: []int{0, 1, 1}, finalCapital: 4},
		{k: 3, w: 0, profits: []int{1, 2, 3}, capital: []int{0, 1, 2}, finalCapital: 6},
		{k: 1, w: 2, profits: []int{1, 2, 3}, capital: []int{1, 1, 2}, finalCapital: 5},
		{k: 2, w: 20, profits: []int{3, 3, 5}, capital: []int{10, 10, 20}, finalCapital: 28},
		{k: 11, w: 11, profits: []int{1, 2, 3}, capital: []int{11, 12, 13}, finalCapital: 17},
		{k: 1, w: 0, profits: []int{1, 2, 3}, capital: []int{1, 1, 2}, finalCapital: 0},
	} {
		finalCapital := findMaximizedCapital(tc.k, tc.w, tc.profits, tc.capital)
		assert.Equal(t, tc.finalCapital, finalCapital)
	}
}
