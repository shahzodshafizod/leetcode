package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestNumEquivDominoPairs$
func TestNumEquivDominoPairs(t *testing.T) {
	for _, tc := range []struct {
		dominoes [][]int
		count    int
	}{
		{dominoes: [][]int{{8, 8}}, count: 0},
		{dominoes: [][]int{{1, 8}}, count: 0},
		{dominoes: [][]int{{8, 1}, {1, 8}}, count: 1},
		{dominoes: [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}, count: 1},
		// {dominoes: [][]int{{1, 1}, {2, 2}, {1, 1}, {1, 2}, {1, 2}, {1, 1}}, count: 4},
		// {dominoes: [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 8}, {1, 9}}, count: 0},
		// {dominoes: [][]int{{1, 8}, {8, 1}, {1, 8}, {1, 8}, {1, 8}, {8, 1}, {8, 2}, {8, 1}, {8, 1}}, count: 28},
		// {dominoes: [][]int{{8, 8}, {8, 8}, {8, 8}, {8, 8}, {8, 8}, {8, 8}, {8, 8}, {8, 8}, {8, 8}}, count: 36},
	} {
		count := numEquivDominoPairs(tc.dominoes)
		assert.Equal(t, tc.count, count)
	}
}
