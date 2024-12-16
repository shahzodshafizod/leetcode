package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestCountPairs$
func TestCountPairs(t *testing.T) {
	for _, tc := range []struct {
		n       int
		edges   [][]int
		queries []int
		answer  []int
	}{
		{n: 4, edges: [][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1}}, queries: []int{2, 3}, answer: []int{6, 5}},
		{n: 5, edges: [][]int{{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5}}, queries: []int{1, 2, 3, 4, 5}, answer: []int{10, 10, 9, 8, 6}},
	} {
		answer := countPairs(tc.n, tc.edges, tc.queries)
		assert.Equal(t, tc.answer, answer)
	}
}
