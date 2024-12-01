package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestShortestDistanceAfterQueries$
func TestShortestDistanceAfterQueries(t *testing.T) {
	for _, tc := range []struct {
		n       int
		queries [][]int
		answer  []int
	}{
		{n: 5, queries: [][]int{{2, 4}, {0, 2}, {0, 4}}, answer: []int{3, 2, 1}},
		{n: 4, queries: [][]int{{0, 3}, {0, 2}}, answer: []int{1, 1}},
		{n: 13, queries: [][]int{{6, 10}, {1, 12}, {7, 9}}, answer: []int{9, 2, 2}},
		{n: 7, queries: [][]int{{1, 4}, {1, 3}, {1, 5}}, answer: []int{4, 4, 3}},
	} {
		answer := shortestDistanceAfterQueries(tc.n, tc.queries)
		assert.Equal(t, tc.answer, answer)
	}
}
